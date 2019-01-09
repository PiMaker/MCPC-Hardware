package interpreter

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"

	"github.com/gdamore/tcell"
	"github.com/mileusna/conditional"
	"github.com/rivo/tview"
)

var symbolMap map[int16]string

func handleError(err error) {
	if err != nil {
		log.Fatalln("ERROR (UART, auto-handled): " + err.Error())
	}
}

// Interpret runs the MCPC debugger
func Interpret(file string, attach bool, maxSteps int, symbolOverride string) {
	var dev *Device

	if attach {
		// Enable hardware debugging mode
		log.Println("attach specified, hardware debugging enabled")

		log.Println("Connecting to device...")
		var err error
		dev, err = establishSerialConnection(file)
		handleError(err)

		log.Println("Enabling debug mode and requesting reset")
		err = dev.setRegister(0, 0)
		handleError(err)
		err = dev.setRegister(0, (1<<0)|(1<<2))
		handleError(err)
		err = dev.setRegister(0, (1 << 0))
		handleError(err)

		log.Println("Waiting for reset...")
		time.Sleep(time.Millisecond * 2000)

		log.Println("Sanity checking")

		regVal, err := dev.getMCPCReg(0xE)
		handleError(err)
		log.Printf("Value read 1: 0x%04x\n", regVal)

		if regVal != 0xFFFF {
			dev.closeConnection()
			log.Fatalln("ERROR: Sanity check 1 failed, is the device connected correctly?")
		}

		regVal, err = dev.getMCPCReg(0xD)
		handleError(err)
		log.Printf("Value read 2: 0x%04x\n", regVal)

		if regVal != 0x0001 {
			dev.closeConnection()
			log.Fatalln("ERROR: Sanity check 2 failed, is the device connected correctly?")
		}

		log.Println("Debugger connection successful!")
	}

	// Read assembly data
	var data []byte

	if attach {
		log.Println("Reading assembly from device (ROM dump)...")

		// Store received data
		for {
			// Read assembly from device
			handleError(dev.triggerROMDump())

			data = make([]byte, 4096)
			checksumLocal := byte(0)
			for insIndex := 0; insIndex < len(data); insIndex += 2 {

				val := make([]byte, 2)
				n, err := dev.port.Read(val)
				handleError(err)

				if n != 2 {
					log.Fatalln("ERROR: Unexpected amount of bytes read")
				}

				data[insIndex] = val[1]
				data[insIndex+1] = val[0]

				checksumLocal ^= val[1]
				checksumLocal ^= val[0]

				fmt.Printf("\r => Read: %d/2048", (insIndex/2)+1)
			}

			// Receive and check checksum
			checksumRemote := make([]byte, 1)
			n, err := dev.port.Read(checksumRemote)
			handleError(err)

			if n != 1 {
				log.Fatalln("ERROR: Unexpected amount of bytes read")
			}

			if checksumLocal != checksumRemote[0] {
				fmt.Printf("\r\n => WARN: Checksum fail, retrying (local: %d != remote: %d)\n", checksumLocal, checksumRemote[0])
				continue
			}

			// Checksum ok, break loop
			break
		}

		// Disable all debugging flags except debug enable
		err := dev.setRegister(0, 1)
		handleError(err)

		fmt.Println("\r\n => ROM dump complete!")

		// Check if device debugger still in sane condition after ROM read
		log.Println("Sanity checking")

		regVal, err := dev.getMCPCReg(0xE)
		handleError(err)
		log.Printf("Value read: 0x%04x\n", regVal)

		if regVal != 0xFFFF {
			dev.closeConnection()
			log.Fatalln("ERROR: Sanity check failed, is the device connected correctly?")
		}

		// Remove trailing 0x0 (HALT) instructions
		log.Println("Trimming HALT trailer")
		for data[len(data)-1] == 0 && data[len(data)-2] == 0 {
			data = data[:len(data)-2]
		}

		if len(data) < 2048 {
			data = append(data, []byte{0, 0}...) // Append single HALT for clarity (if any where trimmed)
		}

	} else {
		log.Println("Reading assembly from file...")
		var err error
		data, err = ioutil.ReadFile(file)
		if err != nil {
			log.Fatalln("ERROR: An error occured reading the input file: " + err.Error())
		}
	}

	// Try to read symbol file
	symbolsFound := false
	if symbolOverride != "" || !attach {
		symbolPath := conditional.String(symbolOverride == "", file+".msym", symbolOverride)
		symbolMap = make(map[int16]string)
		if _, err := os.Stat(symbolPath); err == nil {
			symData, err := ioutil.ReadFile(symbolPath)
			if err != nil {
				log.Fatalln("ERROR: Symbol file found, but an error occured reading it: " + err.Error())
			}

			// Parse msym format
			symSplit := strings.Split(string(symData), ";")
			for _, symEntry := range symSplit {
				if symEntry != "" {
					symEntrySplit := strings.Split(symEntry, "=")
					if len(symEntrySplit) == 2 {
						parsedAddr, err := strconv.ParseInt(symEntrySplit[0], 16, 16)
						if err == nil {
							toAdd := symEntrySplit[1]

							existing, ok := symbolMap[int16(parsedAddr)]
							if ok {
								toAdd = existing + ", " + toAdd
							}

							symbolMap[int16(parsedAddr)] = toAdd
						}
					}
				}
			}

			symbolsFound = true
			log.Println("Symbol file found and loaded!")
		}
	}

	// Parse data into instruction-bounded array
	data16 := make([]uint16, len(data)/2)
	for i := 0; i < len(data16); i++ {
		data16[i] = uint16(data[i*2])<<8 | uint16(data[i*2+1])
	}

	// Run with GUI
	vm := NewVM(data16, 98, 35)

	plength := fmt.Sprintf("0x%04X", len(data16))

	// Set up GUI elements
	root := tview.NewGrid()
	root.SetTitle("MCPC debugger (" + file + ")")
	root.SetRows(4, 18, -3, -1, 2).SetColumns(0, 50)
	root.SetBorder(true)

	cmdField := tview.NewInputField().SetFieldWidth(0).SetLabel("Command: ")
	root.AddItem(cmdField, 4, 0, 1, 1, 0, 0, true)

	disassemblyView := tview.NewTextView()
	disassemblyView.SetBorder(true)
	disassemblyView.SetScrollable(true)
	disassemblyView.SetTitle("Disassembly")
	disassemblyView.SetDynamicColors(true)
	disassemblyView.SetRegions(true)
	root.AddItem(disassemblyView, 0, 0, 4, 1, 0, 0, false)

	// Set up sidebar sections
	stateView := tview.NewTextView()
	stateView.SetBorder(true)
	stateView.SetTitle("VM State")
	stateView.SetText(fmt.Sprintf("State: Not started%s\nPC: 0x0000/%s", conditional.String(symbolsFound, " (msym loaded!)", ""), plength))
	root.AddItem(stateView, 0, 1, 1, 1, 0, 0, false)

	registerView := tview.NewTextView()
	registerView.SetBorder(true)
	registerView.SetTitle("Registers")
	registerView.SetDynamicColors(true)
	registerView.SetRegions(true)
	registerView.SetText(getRegisterText(vm.Registers(), vm.Registers()))
	root.AddItem(registerView, 1, 1, 1, 1, 0, 0, false)

	sramView := tview.NewTable()
	sramView.SetBorder(true)
	sramView.SetTitle("SRAM")
	sramView.SetFixed(1, 0)
	sramView.SetSelectable(true, false)
	sramRow := 0
	sramView.Select(sramRow, 0)
	setSRAMTable(vm, sramView)
	root.AddItem(sramView, 3, 1, 2, 1, 0, 0, false)

	terminalView := tview.NewTextView()
	terminalView.SetBorder(true)
	terminalView.SetTitle("Top of Stack")
	terminalView.SetScrollable(true)
	terminalView.SetDynamicColors(true)
	root.AddItem(terminalView, 2, 1, 1, 1, 0, 0, false)

	terminalText := getStackText(terminalView, vm)
	terminalView.SetText(terminalText)

	virtualPC := vm.Registers().PC.Value

	// Create application
	var modal *tview.Modal
	app := tview.NewApplication()
	app.SetRoot(root, true).SetFocus(root)

	// Set up scrolling
	app.SetInputCapture(func(key *tcell.EventKey) *tcell.EventKey {
		retval := key

		if key.Modifiers() == tcell.ModShift || key.Modifiers() == tcell.ModCtrl {
			if key.Key() == tcell.KeyUp {
				sramRow--
				retval = nil
			} else if key.Key() == tcell.KeyDown {
				sramRow++
				retval = nil
			} else if key.Key() == tcell.KeyPgUp {
				sramRow = 1
				retval = nil
			} else if key.Key() == tcell.KeyPgDn {
				sramRow = sramView.GetRowCount() - 1
				retval = nil
			}

			if sramRow < 0 {
				sramRow = 0
			} else if sramRow >= sramView.GetRowCount() {
				sramRow = sramView.GetRowCount() - 1
			}
			sramView.Select(sramRow, 0)
			app.Draw()
		} else if key.Modifiers() == tcell.ModNone {
			p, _ := strconv.ParseUint(plength[2:], 16, 17)

			if key.Key() == tcell.KeyUp {
				virtualPC--
				retval = nil
			} else if key.Key() == tcell.KeyDown {
				virtualPC++
				retval = nil
			} else if key.Key() == tcell.KeyPgUp {
				virtualPC = 0
				retval = nil
			} else if key.Key() == tcell.KeyPgDn {
				virtualPC = uint16(p) - 1
				retval = nil
			} else if key.Key() == tcell.KeyHome {
				virtualPC = vm.Registers().PC.Value
				retval = nil
			}

			if virtualPC < 0 {
				virtualPC = 0
			} else if virtualPC >= uint16(p) {
				virtualPC = uint16(p) - 1
			}

			disassemblyView.Highlight(fmt.Sprintf("0x%04X", virtualPC))
			disassemblyView.ScrollToHighlight()
			app.Draw()
		}

		return retval
	})

	// Set up behaviours
	cmdField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			cmdField.SetText("")
		} else if key == tcell.KeyEnter {
			// Execute debugger command
			trimmed := strings.TrimSpace(cmdField.GetText())
			split := strings.Split(trimmed, " ")
			if len(split) == 0 {
				split = []string{""}
			}
			switch strings.ToLower(split[0]) {
			case "help":
				messageBox("MCPC Debugger Help", "Arrow keys to move around in disassembly, press HOME to return to current instruction; Available commands: step = Executes a single instruction step (default, use <ENTER> to call with no command in input) :: run <to> = Executes instructions until HALT, BRK or specified PC address <to> is encountered :: runfor <num> = Executes <num> instructions :: exit, quit = Quits the MCPC debugger", app, modal, root)
			case "", "step":
				// Backup values for comparison
				regBck := cloneRegisters(vm.Registers())
				sramBck := make([]uint16, len(vm.SRAM))
				copier.Copy(sramBck, vm.SRAM)
				// Step VM
				_, err := vm.Step()
				// Update view after step
				disassemblyView.SetText(toDisassembly(data16, vm, disassemblyView))
				disassemblyView.Highlight(fmt.Sprintf("0x%04X", vm.Registers().PC.Value))
				virtualPC = vm.Registers().PC.Value
				disassemblyView.ScrollToHighlight()
				if vm.Halted {
					stateView.SetText(fmt.Sprintf("State: Halted\nPC: 0x%04X/%s", vm.Registers().PC.Value, plength))
				} else {
					stateView.SetText(fmt.Sprintf("State: Debugging/Paused\nPC: 0x%04X/%s", vm.Registers().PC.Value, plength))
				}
				registerView.SetText(getRegisterText(vm.Registers(), regBck))
				setSRAMTable(vm, sramView)
				for sramI := 0; sramI < len(sramBck); sramI++ {
					if sramBck[sramI] != vm.SRAM[sramI] {
						sramView.Select(sramI/3+1, 0)
						break
					}
				}
				//terminalText += output
				terminalView.SetText(getStackText(terminalView, vm))
				// Show error message if necessary
				if err != nil {
					messageBox("VM Error", "A VM error occured during the step: "+err.Error(), app, modal, root)
				}

				// Compare with device if running in attached mode
				if attach {
					handleError(dev.step())
					different, diff, err := vm.compareRegistersWithDevice(dev)
					handleError(err)
					if different {
						toShow := "Difference in device registers found:"

						for _, d := range diff {
							toShow += " [" + d + "]"
						}

						toShow += " - State is now invalid"

						messageBox("Device inconsistency", toShow, app, modal, root)
					}
				}

			case "run", "runfor":
				// Run until BRK, timeout or match
				stateView.SetText(fmt.Sprintf("State: Running\nPC: -"))
				app.Draw()

				timeout := 0
				maxTimeout := 10000
				match := -1
				if len(split) > 1 {
					if split[0] == "run" {
						m, cerr := strconv.ParseInt(split[1], 16, 17)
						if cerr == nil {
							match = int(m)
						} else {
							m, cerr = strconv.ParseInt(split[1][2:], 16, 17)
							if cerr == nil {
								match = int(m)
							} else {
								messageBox("Warning", "You passed a parameter to run, however it could not be parsed as a hex number. It will be ignored.", app, modal, root)
							}
						}
					} else if split[0] == "runfor" {
						m, cerr := strconv.ParseInt(split[1], 10, 32)
						if cerr == nil {
							maxTimeout = int(m)
						} else {
							messageBox("Warning", "You passed a parameter to runfor, however it could not be parsed as a number. It will be ignored.", app, modal, root)
						}
					}
				}

				for !vm.Halted && timeout < maxTimeout {
					brk, err := vm.Step()
					if err != nil {
						messageBox("VM Error", fmt.Sprintf("A VM error occured during step 0x%X (at PC=0x%X): %s", vm.EEPROM[vm.Registers().PC.Value-1], vm.Registers().PC.Value-1, err.Error()), app, modal, root)
					}

					// Compare with device if running in attached mode
					if attach {
						handleError(dev.step())
						different, diff, err := vm.compareRegistersWithDevice(dev)
						handleError(err)
						if different {
							toShow := "Difference in device registers found:"

							for _, d := range diff {
								toShow += " [" + d + "]"
							}

							toShow += " - State is now invalid"

							messageBox("Device inconsistency", toShow, app, modal, root)

							break
						}
					}

					if (match == -1 && brk && split[0] == "run") || int(vm.Registers().PC.Value) == match {
						break
					}
					timeout++
				}

				if timeout == 10000 && split[0] == "run" {
					messageBox("Timeout", "Execution paused because a timeout was reached (10000 steps).", app, modal, root)
				}

				// Update view after steps
				disassemblyView.SetText(toDisassembly(data16, vm, disassemblyView))
				disassemblyView.Highlight(fmt.Sprintf("0x%04X", vm.Registers().PC.Value))
				virtualPC = vm.Registers().PC.Value
				disassemblyView.ScrollToHighlight()
				if vm.Halted {
					stateView.SetText(fmt.Sprintf("State: Halted\nPC: 0x%04X/%s", vm.Registers().PC.Value, plength))
				} else {
					stateView.SetText(fmt.Sprintf("State: Debugging/Paused\nPC: 0x%04X/%s", vm.Registers().PC.Value, plength))
				}
				registerView.SetText(getRegisterText(vm.Registers(), vm.Registers()))
				setSRAMTable(vm, sramView)
				terminalView.SetText(getStackText(terminalView, vm))
			case "exit", "quit":
				app.Stop()
			default:
				messageBox("Invalid command", "Type \"help\" to see a list of available commands.", app, modal, root)
			}

			cmdField.SetText("")
		}
	})

	// Set disassembly text last to avoid width glitching for label offsets
	// Nevermind, doesn't work either way
	disassemblyView.SetText(toDisassembly(data16, vm, disassemblyView))
	disassemblyView.Highlight(fmt.Sprintf("0x%04X", virtualPC))

	// Run GUI app

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}

	if attach && dev != nil {
		dev.closeConnection()
	}
}

func messageBox(title, text string, app *tview.Application, modal *tview.Modal, root *tview.Grid) {
	m := tview.NewModal()
	m.SetBorder(true)
	m.SetTitle(title)
	m.SetText(text)
	m.AddButtons([]string{"Ok"})
	focused := app.GetFocus()
	m.SetDoneFunc(func(_ int, _ string) {
		app.SetRoot(root, true)
		app.SetFocus(focused)
		modal = nil
	})
	modal = m
	app.SetRoot(modal, false)
	app.SetFocus(modal)
}

func getRegisterText(reg, regOld *Registers) string {
	retval := ""
	retval += "[gray]0x0[" + colorRegister + "]    A[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.A.Value == regOld.A.Value, "", "[red]"), reg.A.Value, fmt.Sprintf(" <> 0x%04X", regOld.A.Value))
	retval += "[gray]0x1[" + colorRegister + "]    B[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.B.Value == regOld.B.Value, "", "[red]"), reg.B.Value, fmt.Sprintf(" <> 0x%04X", regOld.B.Value))
	retval += "[gray]0x2[" + colorRegister + "]    C[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.C.Value == regOld.C.Value, "", "[red]"), reg.C.Value, fmt.Sprintf(" <> 0x%04X", regOld.C.Value))
	retval += "[gray]0x3[" + colorRegister + "]    D[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.D.Value == regOld.D.Value, "", "[red]"), reg.D.Value, fmt.Sprintf(" <> 0x%04X", regOld.D.Value))
	retval += "[gray]0x4[" + colorRegister + "]    E[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.E.Value == regOld.E.Value, "", "[red]"), reg.E.Value, fmt.Sprintf(" <> 0x%04X", regOld.E.Value))
	retval += "[gray]0x5[" + colorRegister + "]    F[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.F.Value == regOld.F.Value, "", "[red]"), reg.F.Value, fmt.Sprintf(" <> 0x%04X", regOld.F.Value))
	retval += "[gray]0x6[" + colorRegister + "]    G[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.G.Value == regOld.G.Value, "", "[red]"), reg.G.Value, fmt.Sprintf(" <> 0x%04X", regOld.G.Value))
	retval += "[gray]0x7[" + colorRegister + "]    H[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.H.Value == regOld.H.Value, "", "[red]"), reg.H.Value, fmt.Sprintf(" <> 0x%04X", regOld.H.Value))
	retval += "[gray]0x8[" + colorRegister + "] SCR1[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.SCR1.Value == regOld.SCR1.Value, "", "[red]"), reg.SCR1.Value, fmt.Sprintf(" <> 0x%04X", regOld.SCR1.Value))
	retval += "[gray]0x9[" + colorRegister + "] SCR2[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.SCR2.Value == regOld.SCR2.Value, "", "[red]"), reg.SCR2.Value, fmt.Sprintf(" <> 0x%04X", regOld.SCR2.Value))
	retval += "[gray]0xA[" + colorRegister + "]   SP[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.SP.Value == regOld.SP.Value, "", "[red]"), reg.SP.Value, fmt.Sprintf(" <> 0x%04X", regOld.SP.Value))
	retval += "[gray]0xB[" + colorRegister + "]   PC[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.PC.Value == regOld.PC.Value+1, "", "[red]"), reg.PC.Value, fmt.Sprintf(" <> 0x%04X", regOld.PC.Value))
	retval += "[gray]0xC[" + colorRegister + "]    0[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.Zero.Value == regOld.Zero.Value, "", "[red]"), reg.Zero.Value, fmt.Sprintf(" <> 0x%04X", regOld.Zero.Value))
	retval += "[gray]0xD[" + colorRegister + "]    1[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.One.Value == regOld.One.Value, "", "[red]"), reg.One.Value, fmt.Sprintf(" <> 0x%04X", regOld.One.Value))
	retval += "[gray]0xE[" + colorRegister + "]   -1[white] = " + fmt.Sprintf("%s0x%04X%s\n", conditional.String(reg.NegOne.Value == regOld.NegOne.Value, "", "[red]"), reg.NegOne.Value, fmt.Sprintf(" <> 0x%04X", regOld.NegOne.Value))
	retval += "[gray]0xF[" + colorRegister + "]  BUS[white] = " + fmt.Sprintf("%s0x%04X%s", conditional.String(reg.BUS.Value == regOld.BUS.Value, "", "[red]"), reg.BUS.Value, fmt.Sprintf(" <> 0x%04X", regOld.BUS.Value))
	return retval
}

const colorNotes = "red"
const colorPCAddr = "gray"
const colorRawIns = "lightgray"
const colorCmd = "lightgreen"
const colorRegister = "yellow"

var formatRemoverRegex = regexp.MustCompile(`\[.*?\]`)

func toDisassembly(raw []uint16, vm *VM, view *tview.TextView) string {
	retval := ""
	skip := false

	for i, ins := range raw {
		addr := fmt.Sprintf("0x%04X", i)

		if skip {
			skip = false
			retval += fmt.Sprintf("[\"%s\"][%s]%s  ...\n", addr, colorPCAddr, addr)
			continue
		}

		cmd, params, note, set := decodeAssembly(ins, vm)
		if set {
			if len(raw)-1 == i {
				log.Fatalln("ERROR: SET is last instruction")
			}

			params += fmt.Sprintf("0x%04X", raw[i+1])
			skip = true
		}

		tabs := " "
		for t := 0; t < 5-len(cmd); t++ {
			tabs += " "
		}
		retval += fmt.Sprintf("[\"%s\"][%s]%s  [%s]0x%04X  [%s]%s%s[white]%s", addr, colorPCAddr, addr, colorRawIns, ins, colorCmd, cmd, tabs, params)

		if note != "" {
			if params != "" {
				retval += " :: "
			}
			retval += fmt.Sprintf("[%s](%s)[white]", colorNotes, note)
		}

		label, ok := symbolMap[int16(i)]
		if ok {
			// Symbol found, append right justified
			_, _, width, _ := view.GetRect()

			// Get length of current string
			formatRemovedLines := strings.Split(formatRemoverRegex.ReplaceAllString(retval, ""), "\n")
			curWidth := len(formatRemovedLines[len(formatRemovedLines)-1])

			// Calculate offset
			offset := width - curWidth - len(label)

			// Print label with prepending format spaces for right justification
			for k := 0; k < offset-3; k++ {
				retval += " "
			}

			retval += "[blue]" + label + "[white]"
		}

		retval += "[\"\"]\n"
	}

	return retval[:len(retval)-2]
}

var sramWriteWaitingDecoder = false

func decodeAssembly(c uint16, vm *VM) (cmd, params, note string, set bool) {
	set = false

	switch c & 0x000F {
	case 0x0:
		cmd = "HALT"
	case 0x1:
		valueToMove := GetReg(vm, c, regFrom).Value

		if GetReg(vm, c, regTo) == vm.Registers().PC {
			cmd = "JMP"
			params = fmt.Sprintf("to 0x%04x", valueToMove)

			label, ok := symbolMap[int16(valueToMove)]
			if ok {
				note = fmt.Sprintf("label: [blue]%s[%s]", label, colorNotes)
			}

			break
		}

		cmd = "MOV"
		params = decodeRegister(byte((c&0x00F0)>>4), "white") + " -> " + decodeRegister(byte((c&0x0F00)>>8), "white")
		note = fmt.Sprintf("set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), valueToMove)
	case 0x2:
		cmd = "MOVNZ"
		valueToMove := GetReg(vm, c, regFrom).Value

		if GetReg(vm, c, regTo) == vm.Registers().PC {
			cmd = "JMPNZ"
		}

		params = decodeRegister(byte((c&0x00F0)>>4), "white") + " -> " + decodeRegister(byte((c&0x0F00)>>8), "white") + " if " + decodeRegister(byte((c&0xF000)>>12), "white") + " != 0"
		if GetReg(vm, c, regIf).Value != 0 {
			note = fmt.Sprintf("TRUE: set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), valueToMove)
		} else {
			note = fmt.Sprintf("FALSE, would do: set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), valueToMove)
		}

		label, ok := symbolMap[int16(valueToMove)]
		if ok {
			note += fmt.Sprintf("; label: [blue]%s[%s]", label, colorNotes)
		}
	case 0x3:
		cmd = "MOVEZ"
		valueToMove := GetReg(vm, c, regFrom).Value

		if GetReg(vm, c, regTo) == vm.Registers().PC {
			cmd = "JMPEZ"
		}

		params = decodeRegister(byte((c&0x00F0)>>4), "white") + " -> " + decodeRegister(byte((c&0x0F00)>>8), "white") + " if " + decodeRegister(byte((c&0xF000)>>12), "white") + " == 0"
		if GetReg(vm, c, regIf).Value == 0 {
			note = fmt.Sprintf("TRUE: set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), valueToMove)
		} else {
			note = fmt.Sprintf("FALSE, would do: set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), valueToMove)
		}

		label, ok := symbolMap[int16(valueToMove)]
		if ok {
			note += fmt.Sprintf("; label: [blue]%s[%s]", label, colorNotes)
		}
	case 0x4:
		cmd = "BUS"
		params = ""
		note = "Deprecated!"
	case 0x5:
		cmd = "MEMR"

		if GetReg(vm, c, regFrom) == vm.Registers().SP {
			cmd = "POP"
		}

		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " <- @" + decodeRegister(byte((c&0x00F0)>>4), "white")
		addr := GetReg(vm, c, regFrom).Value

		if (addr & 0x8000) == 0 {
			note = fmt.Sprintf("Read data @%04x (=%04x) into register %s", addr, vm.SRAM[addr], decodeRegister(byte((c&0x0F00)>>8), colorNotes))
		} else if addr == 0x8000 {
			note = fmt.Sprintf("Read data @%04x (MCPC version = 0x8001 [VM]) into register %s", addr, decodeRegister(byte((c&0x0F00)>>8), colorNotes))
		} else if addr >= 0xD000 && addr < 0xD800 {
			note = fmt.Sprintf("Read ROM-data @%04x (ROM @%04x) (=%04x) into register %s", addr, addr-0xD000, vm.EEPROM[addr-0xD000], decodeRegister(byte((c&0x0F00)>>8), colorNotes))
		} else {
			note = "STUB: Read from unknown CFG"
		}
	case 0x6:
		cmd = "SET"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " to "
		set = true
	case 0x7:
		cmd = "MEMW"

		if GetReg(vm, c, regFrom) == vm.Registers().SP {
			cmd = "PUSH"
		}

		params = decodeRegister(byte((c&0xF000)>>12), "white") + " -> @" + decodeRegister(byte((c&0x00F0)>>4), "white")
		addr := GetReg(vm, c, regFrom).Value

		if (addr & 0x8000) == 0 {
			note = fmt.Sprintf("Write data from register %s (=%04x) into RAM @%04x", decodeRegister(byte((c&0x0F00)>>8), colorNotes), GetReg(vm, c, regIf).Value, addr)
		} else {
			note = "STUB: Write to unknown CFG"
		}
	case 0x8:
		cmd = "AND"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " & " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = fmt.Sprintf("%04X & %04X = %04X", GetReg(vm, c, regFrom).Value, GetReg(vm, c, regOp).Value, (GetReg(vm, c, regFrom).Value & GetReg(vm, c, regOp).Value))
	case 0x9:
		cmd = "OR"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " | " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = fmt.Sprintf("%04X | %04X = %04X", GetReg(vm, c, regFrom).Value, GetReg(vm, c, regOp).Value, (GetReg(vm, c, regFrom).Value | GetReg(vm, c, regOp).Value))
	case 0xA:
		cmd = "XOR"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " ^ " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = fmt.Sprintf("%04X ^ %04X = %04X", GetReg(vm, c, regFrom).Value, GetReg(vm, c, regOp).Value, (GetReg(vm, c, regFrom).Value ^ GetReg(vm, c, regOp).Value))
		if GetReg(vm, c, regFrom).Value == 0xFFFF || GetReg(vm, c, regOp).Value == 0xFFFF {
			note += " (COM)"
		}
	case 0xB:
		cmd = "ADD"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " + " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = fmt.Sprintf("%04X + %04X = %04X", GetReg(vm, c, regFrom).Value, GetReg(vm, c, regOp).Value, (GetReg(vm, c, regFrom).Value + GetReg(vm, c, regOp).Value))
	case 0xC:
		cmd = "SHFT"
		byVal := GetReg(vm, c, regOp).Value
		if byVal&0xFF00 == 0 {
			params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " >> " + decodeRegister(byte((c&0xF000)>>12), "white")
			note = fmt.Sprintf("%04X >> %02X = %04X (dir: %02X == 0, right)", GetReg(vm, c, regFrom).Value, byVal&0x00FF, (GetReg(vm, c, regFrom).Value >> (byVal & 0x00FF)), byVal>>8)
		} else {
			params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " << " + decodeRegister(byte((c&0xF000)>>12), "white")
			note = fmt.Sprintf("%04X << %02X = %04X (dir: %02X != 0, left)", GetReg(vm, c, regFrom).Value, byVal&0x00FF, (GetReg(vm, c, regFrom).Value << (byVal & 0x00FF)), byVal>>8)
		}

	case 0xD:
		cmd = "MUL"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " * " + decodeRegister(byte((c&0xF000)>>12), "white")
		mulRes := int(GetReg(vm, c, regFrom).Value) * int(GetReg(vm, c, regOp).Value)
		note = fmt.Sprintf("%04X * %04X = %04X (Overflow: %t)", GetReg(vm, c, regFrom).Value, GetReg(vm, c, regOp).Value, GetReg(vm, c, regFrom).Value*GetReg(vm, c, regOp).Value, mulRes > 0xFFFF)
	case 0xE:
		cmd = "GT"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " > " + decodeRegister(byte((c&0xF000)>>12), "white")
		val1 := GetReg(vm, c, regFrom).Value
		val2 := GetReg(vm, c, regOp).Value
		if val1 > val2 {
			note = fmt.Sprintf("%04X > %04X = 0xFFFF", val1, val2)
		} else {
			note = fmt.Sprintf("%04X > %04X = 0x0", val1, val2)
		}
	case 0xF:
		cmd = "EQ"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " == " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = "t/f = 0x0/0xFFFF"
		val1 := GetReg(vm, c, regFrom).Value
		val2 := GetReg(vm, c, regOp).Value
		if val1 == val2 {
			note = fmt.Sprintf("%04X == %04X = 0xFFFF", val1, val2)
		} else {
			note = fmt.Sprintf("%04X == %04X = 0x0", val1, val2)
		}
	}

	return cmd, params, note, set
}

func decodeRegister(reg byte, origColor string) string {
	retval := "[" + colorRegister + "]"
	switch reg {
	case 0x0:
		retval += "A"
	case 0x1:
		retval += "B"
	case 0x2:
		retval += "C"
	case 0x3:
		retval += "D"
	case 0x4:
		retval += "E"
	case 0x5:
		retval += "F"
	case 0x6:
		retval += "G"
	case 0x7:
		retval += "H"
	case 0x8:
		retval += "SCR1"
	case 0x9:
		retval += "SCR2"
	case 0xA:
		retval += "SP"
	case 0xB:
		retval += "PC"
	case 0xC:
		retval += "0(r)"
	case 0xD:
		retval += "+1(r)"
	case 0xE:
		retval += "-1(r)"
	case 0xF:
		retval += "BUS(r)"
	default:
		retval += "INVALID_REGISTER"
	}

	return retval + "[" + origColor + "]"
}

func cloneRegisters(reg *Registers) *Registers {
	retval := &Registers{
		A:      &Register{},
		B:      &Register{},
		C:      &Register{},
		D:      &Register{},
		E:      &Register{},
		F:      &Register{},
		G:      &Register{},
		H:      &Register{},
		SCR1:   &Register{},
		SCR2:   &Register{},
		SP:     &Register{},
		PC:     &Register{},
		Zero:   &Register{},
		One:    &Register{},
		NegOne: &Register{},
		BUS:    &Register{},
	}
	copier.Copy(retval.A, reg.A)
	copier.Copy(retval.B, reg.B)
	copier.Copy(retval.C, reg.C)
	copier.Copy(retval.D, reg.D)
	copier.Copy(retval.E, reg.E)
	copier.Copy(retval.F, reg.F)
	copier.Copy(retval.G, reg.G)
	copier.Copy(retval.H, reg.H)
	copier.Copy(retval.SCR1, reg.SCR1)
	copier.Copy(retval.SCR2, reg.SCR2)
	copier.Copy(retval.SP, reg.SP)
	copier.Copy(retval.PC, reg.PC)
	copier.Copy(retval.Zero, reg.Zero)
	copier.Copy(retval.One, reg.One)
	copier.Copy(retval.NegOne, reg.NegOne)
	copier.Copy(retval.BUS, reg.BUS)
	return retval
}

func setSRAMTable(vm *VM, tbl *tview.Table) {
	// Header
	tbl.SetCell(0, 0, &tview.TableCell{
		Text:  "Addr",
		Color: tcell.ColorRed,
	})
	tbl.SetCell(0, 1, &tview.TableCell{
		Text:  "Value",
		Color: tcell.ColorRed,
	})
	tbl.SetCell(0, 2, &tview.TableCell{
		Text:  "Addr",
		Color: tcell.ColorRed,
	})
	tbl.SetCell(0, 3, &tview.TableCell{
		Text:  "Value",
		Color: tcell.ColorRed,
	})
	tbl.SetCell(0, 4, &tview.TableCell{
		Text:  "Addr",
		Color: tcell.ColorRed,
	})
	tbl.SetCell(0, 5, &tview.TableCell{
		Text:  "Value",
		Color: tcell.ColorRed,
	})

	// Value
	for i := uint16(0); i <= MaxSRAMValue+3; i += 3 {
		for c := uint16(0); c < 3; c++ {
			if int(i+c) >= len(vm.SRAM) {
				break
			}

			tbl.SetCell(int(i)/3+1, int(c)*2, &tview.TableCell{
				Text:  fmt.Sprintf("0x%04X", i+c),
				Color: tcell.ColorGray,
			})
			tbl.SetCell(int(i)/3+1, int(c)*2+1, &tview.TableCell{
				Text:  fmt.Sprintf("0x%04X", vm.SRAM[i+c]),
				Color: tcell.ColorWhite,
			})
		}
	}
}

func getStackText(view *tview.TextView, vm *VM) string {
	_, _, _, lines := view.GetInnerRect()
	addr := vm.Registers().SP.Value
	retval := ""

	if addr > 0 && addr < 0x7FFF {
		for i := 0; i < lines; i++ {
			if addr == 0x7FFF {
				break
			}

			retval = fmt.Sprintf("%s[gray]0x%04x[white] 0x%04x%s\n", retval, addr, vm.SRAM[addr], conditional.String(i == 0, " @SP", ""))
			addr++
		}
	} else if addr == 0x7FFF {
		return "Stack empty"
	} else {
		return fmt.Sprintf("SP invalid: 0x%04x", addr)
	}

	return retval
}
