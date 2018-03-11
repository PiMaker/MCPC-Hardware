package interpreter

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jinzhu/copier"

	"github.com/gdamore/tcell"
	"github.com/mileusna/conditional"
	"github.com/rivo/tview"
)

// Interpret runs a .mb binary on a virtual machine
func Interpret(file, config string, gui bool) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("ERROR: An error occured reading the input file.")
	}

	// Parse data into instruction-bounded array
	data16 := make([]uint16, len(data)/2)
	for i := 0; i < len(data16); i++ {
		data16[i] = uint16(data[i*2])<<8 | uint16(data[i*2+1])
	}

	if gui {

		// Run with GUI
		vm := NewVM(data16)

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
		disassemblyView.SetText(toDisassembly(data16, vm))
		disassemblyView.Highlight(fmt.Sprintf("0x%04X", vm.Registers.PC.Value))
		root.AddItem(disassemblyView, 0, 0, 4, 1, 0, 0, false)

		// Set up sidebar sections
		stateView := tview.NewTextView()
		stateView.SetBorder(true)
		stateView.SetTitle("VM State")
		stateView.SetText("State: Not started\nPC: 0x0000/" + plength)
		root.AddItem(stateView, 0, 1, 1, 1, 0, 0, false)

		registerView := tview.NewTextView()
		registerView.SetBorder(true)
		registerView.SetTitle("Registers")
		registerView.SetDynamicColors(true)
		registerView.SetRegions(true)
		registerView.SetText(getRegisterText(vm.Registers, vm.Registers))
		root.AddItem(registerView, 1, 1, 1, 1, 0, 0, false)

		sramView := tview.NewTable()
		sramView.SetBorder(true)
		sramView.SetTitle("SRAM")
		root.AddItem(sramView, 2, 1, 1, 1, 0, 0, false)

		terminalView := tview.NewTextView()
		terminalView.SetBorder(true)
		terminalView.SetTitle("Terminal output")
		terminalView.SetScrollable(true)
		terminalText := ""
		terminalView.SetText(terminalText)
		root.AddItem(terminalView, 3, 1, 2, 1, 0, 0, false)

		// Create application
		var modal *tview.Modal
		app := tview.NewApplication()
		app.SetRoot(root, true).SetFocus(root)

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
				switch split[0] {
				case "", "step":
					// Backup register values for comparison
					regBck := cloneRegisters(vm.Registers)
					// Step VM
					_, output, err := vm.Step()
					// Update view after step
					disassemblyView.SetText(toDisassembly(data16, vm))
					disassemblyView.Highlight(fmt.Sprintf("0x%04X", vm.Registers.PC.Value))
					disassemblyView.ScrollToHighlight()
					stateView.SetText(fmt.Sprintf("State: Debugging/Paused\nPC: 0x%04X/%s", vm.Registers.PC.Value, plength))
					registerView.SetText(getRegisterText(vm.Registers, regBck))
					terminalText += output
					terminalView.SetText(strings.Replace(terminalText, "\n", "\\n\n", -1))
					terminalView.ScrollToEnd()
					// Show error message if necessary
					if err != nil {
						messageBox("Error", "A VM error occured during the step: "+err.Error(), app, modal, root)
					}
				default:
					messageBox("Invalid command", "Type \"help\" to see a list of available commands.", app, modal, root)
				}

				cmdField.SetText("")
			}
		})

		// Run GUI app

		if err := app.Run(); err != nil {
			log.Fatalln(err)
		}

	} else {

		// GUI-less interpretation
		log.Println("Interpreting " + file + "...")

		// Create and run VM
		vm := NewVM(data16)
		for !vm.Halted {
			// BRK is ignored
			_, termout, err := vm.Step()
			if err != nil {
				log.Fatalf("ERROR on instruction 0x%X (at PC=0x%X): %s\n", vm.EEPROM[vm.Registers.PC.Value-1], vm.Registers.PC.Value-1, err.Error())
			}
			if termout != "" {
				fmt.Print(termout)
			}
		}

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

func toDisassembly(raw []uint16, vm *VM) string {
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
				retval += "\t"
			}
			retval += fmt.Sprintf("[%s](%s)[white]", colorNotes, note)
		}

		retval += "[\"\"]\n"
	}

	return retval[:len(retval)-2]
}

func decodeAssembly(c uint16, vm *VM) (cmd, params, note string, set bool) {
	set = false

	switch c & 0x000F {
	case 0x0:
		cmd = "HALT"
	case 0x1:
		cmd = "MOV"
		params = decodeRegister(byte((c&0x00F0)>>4), "white") + " -> " + decodeRegister(byte((c&0x0F00)>>8), "white")
		note = fmt.Sprintf("set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), getReg(vm, c, regFrom).Value)
	case 0x2:
		cmd = "MOVNZ"
		params = decodeRegister(byte((c&0x00F0)>>4), "white") + " -> " + decodeRegister(byte((c&0x0F00)>>8), "white") + " if " + decodeRegister(byte((c&0xF000)>>12), "white") + " != 0"
		if getReg(vm, c, regFrom).Value != 0 {
			note = fmt.Sprintf("TRUE: set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), getReg(vm, c, regFrom).Value)
		} else {
			note = fmt.Sprintf("FALSE, would do: set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), getReg(vm, c, regFrom).Value)
		}
	case 0x3:
		cmd = "MOVEZ"
		params = decodeRegister(byte((c&0x00F0)>>4), "white") + " -> " + decodeRegister(byte((c&0x0F00)>>8), "white") + " if " + decodeRegister(byte((c&0xF000)>>12), "white") + " == 0"
		if getReg(vm, c, regFrom).Value == 0 {
			note = fmt.Sprintf("TRUE: set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), getReg(vm, c, regFrom).Value)
		} else {
			note = fmt.Sprintf("FALSE, would do: set %s to 0x%04X", decodeRegister(byte((c&0x0F00)>>8), colorNotes), getReg(vm, c, regFrom).Value)
		}
	case 0x4:
		cmd = "BUS"
		params = "send " + decodeRegister(byte((c&0x00F0)>>4), "white") + " to " + fmt.Sprintf("0x%02X", byte((c&0x0F00)>>8)) + " on bus"
		note = fmt.Sprintf("payload=%04X", getReg(vm, c, regFrom).Value)
	case 0x5:
		cmd = "HOLD"
	case 0x6:
		cmd = "SET"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " to "
		set = true
	case 0x7:
		cmd = "BRK"
	case 0x8:
		cmd = "AND"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " & " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = fmt.Sprintf("%04X & %04X = %04X", getReg(vm, c, regFrom).Value, getReg(vm, c, regOp).Value, (getReg(vm, c, regFrom).Value & getReg(vm, c, regOp).Value))
	case 0x9:
		cmd = "OR"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " | " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = fmt.Sprintf("%04X | %04X = %04X", getReg(vm, c, regFrom).Value, getReg(vm, c, regOp).Value, (getReg(vm, c, regFrom).Value | getReg(vm, c, regOp).Value))
	case 0xA:
		cmd = "NOT"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = !" + decodeRegister(byte((c&0x00F0)>>4), "white")
		note = fmt.Sprintf("!%04X = %04X", getReg(vm, c, regFrom).Value, ^getReg(vm, c, regFrom).Value)
	case 0xB:
		cmd = "ADD"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " + " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = fmt.Sprintf("%04X + %04X = %04X", getReg(vm, c, regFrom).Value, getReg(vm, c, regOp).Value, (getReg(vm, c, regFrom).Value + getReg(vm, c, regOp).Value))
	case 0xC:
		cmd = "SHFT"
		sval := byte((c & 0xF000) >> 12)
		if sval&0x8 == 0 {
			params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " >> 0x" + fmt.Sprintf("%X", sval)
			note = fmt.Sprintf("= %04X", getReg(vm, c, regFrom).Value>>sval)
		} else {
			params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " << 0x" + fmt.Sprintf("%X", sval&0x7)
			note = fmt.Sprintf("= %04X (negative shift)", getReg(vm, c, regFrom).Value<<(sval&0x7))
		}
	case 0xD:
		cmd = "MUL"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " * " + decodeRegister(byte((c&0xF000)>>12), "white")
		mulRes := int(getReg(vm, c, regFrom).Value) * int(getReg(vm, c, regOp).Value)
		note = fmt.Sprintf("%04X * %04X = %04X (Overflow: %t)", getReg(vm, c, regFrom).Value, getReg(vm, c, regOp).Value, getReg(vm, c, regFrom).Value*getReg(vm, c, regOp).Value, mulRes > 0xFFFF)
	case 0xE:
		cmd = "GT"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " > " + decodeRegister(byte((c&0xF000)>>12), "white")
		val1 := getReg(vm, c, regFrom).Value
		val2 := getReg(vm, c, regOp).Value
		if val1 > val2 {
			note = fmt.Sprintf("%04X > %04X = 0xFFFF", val1, val2)
		} else {
			note = fmt.Sprintf("%04X > %04X = 0x0", val1, val2)
		}
	case 0xF:
		cmd = "EQ"
		params = decodeRegister(byte((c&0x0F00)>>8), "white") + " = " + decodeRegister(byte((c&0x00F0)>>4), "white") + " == " + decodeRegister(byte((c&0xF000)>>12), "white")
		note = "t/f = 0x0/0xFFFF"
		val1 := getReg(vm, c, regFrom).Value
		val2 := getReg(vm, c, regOp).Value
		if val1 == val2 {
			note = fmt.Sprintf("%04X = %04X = 0xFFFF", val1, val2)
		} else {
			note = fmt.Sprintf("%04X = %04X = 0x0", val1, val2)
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
