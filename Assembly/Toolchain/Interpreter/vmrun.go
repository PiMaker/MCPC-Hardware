package interpreter

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"unicode"

	"github.com/nsf/termbox-go"
)

var (
	keycodeLookupRunes = map[rune]uint32{
		'a': 0x1c,
	}

	keycodeLookupTermbox = map[termbox.Key]uint32{
		termbox.KeyBackspace: 0x66,
	}

	keycodeBreak   = uint32(0x00F00000)
	keycodeLSHFT   = uint32(0x00120000)
	keyboardIrqNum = uint32(0xA)
)

// VMRun executes the given file in a virtual MCPC
func VMRun(file string) {

	// Termbox init
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.HideCursor()
	termbox.Sync()

	// Retrieve term info
	width, height := termbox.Size()

	// Draw border
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == 0 || x == width-1 || (x == 121 && y <= 67 && y > 0 && y < height-3) {
				// Left/Right border
				termbox.SetCell(x, y, '|', termbox.ColorWhite, termbox.ColorBlack)
			}

			if y == 0 || y == height-1 || y == height-3 || (y == 66 && height > 67 && x <= 121 && x > 0) {
				// Top/Bottom/State border
				termbox.SetCell(x, y, '-', termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	// Set manual border nubs
	termbox.SetCell(0, 0, '+', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(width-1, 0, '+', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(0, height-1, '+', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(width-1, height-1, '+', termbox.ColorWhite, termbox.ColorBlack)

	termbox.SetCell(0, height-3, '+', termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(width-1, height-3, '+', termbox.ColorWhite, termbox.ColorBlack)

	if width > 121 {
		termbox.SetCell(121, 0, '+', termbox.ColorWhite, termbox.ColorBlack)
	}
	if height > 67 {
		termbox.SetCell(0, 66, '+', termbox.ColorWhite, termbox.ColorBlack)
	}

	termbox.Flush()

	// Event loop (goroutine)
	closeChan := make(chan bool, 1)
	irqChan := make(chan uint32, 256)
	go func() {
		for {
			event := termbox.PollEvent()

			if event.Type == termbox.EventInterrupt {
				closeChan <- true
			}

			if event.Type == termbox.EventKey && event.Ch == 0 && event.Key == termbox.KeyCtrlC {
				closeChan <- true
			}

			if event.Type == termbox.EventKey {
				// Send keyboard irq
				if event.Ch == 0 {
					// Special key
					keyCode, ok := keycodeLookupTermbox[event.Key]
					keyCode = keyCode << 16
					if ok {
						// Send MAKE
						irqChan <- keyCode | keyboardIrqNum

						// Send BREAK
						irqChan <- keycodeBreak | keyboardIrqNum
						irqChan <- keyCode | keyboardIrqNum
					}
				} else {
					// Character key
					letter := event.Ch

					uppercase := unicode.IsUpper(letter)
					if uppercase {
						// Uppercase letter, send shift first
						irqChan <- keycodeLSHFT | keyboardIrqNum
						letter = unicode.ToLower(letter)
					}

					keyCode, ok := keycodeLookupRunes[letter]
					keyCode = keyCode << 16
					if ok {
						// Send MAKE
						irqChan <- keyCode | keyboardIrqNum

						// Send BREAK
						irqChan <- keycodeBreak | keyboardIrqNum
						irqChan <- keyCode | keyboardIrqNum
					}

					// Send SHFT BREAK in case of uppercase
					if uppercase {
						irqChan <- keycodeBreak | keyboardIrqNum
						irqChan <- keycodeLSHFT | keyboardIrqNum
					}
				}
			}
		}
	}()

	// Load data from file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		termbox.Close()
		log.Fatalln("ERROR: An error occured reading the input file: " + err.Error())
	}

	// Parse data into instruction-bounded array
	data16 := make([]uint16, len(data)/2)
	for i := 0; i < len(data16); i++ {
		data16[i] = uint16(data[i*2])<<8 | uint16(data[i*2+1])
	}

	// VM init
	vm := NewVM(data16, uint16(width-2), uint16(height-4))
	writeVMState(vm, height, -1)

	// Attach VGA update handler
	vgaChanged := false
	vm.VgaChangeCallback = func(addr, x, y, old, new uint16) {
		r := rune(new & 0x00FF)

		// Treat special characters like spaces/blanks
		if r < ' ' {
			r = ' '
		}

		// Only update terminal if output to display has actually changed
		buf := termbox.CellBuffer()
		if buf[(int(y)+1)*width+(int(x)+1)].Ch != r {
			termbox.SetCell(int(x)+1, int(y)+1, r, termbox.ColorWhite, termbox.ColorBlack)
			vgaChanged = true
		}
	}

	// Speed counter
	stepCounter := 0
	lastSecond := -1
	currentSpeed := -1

	// Event loop
	for {
		// Inject IRQ if needed
		for len(irqChan) > 0 {
			vm.InjectIRQ(<-irqChan)
		}

		flushTerminal := false

		if !vm.Halted {
			// Check for changes in tracked states
			preH := vm.RegDef.H.Value
			preHirq := vm.RegIrq.H.Value
			preIrqEn := vm.IrqEn
			preInIrq := vm.InIrq

			// Perform VM step
			// TODO: Implement debug BRK?
			_, err := vm.Step()

			if err != nil {
				termbox.Close()
				log.Fatalln("VM ERROR: " + err.Error())
			}

			// Changecheck 2
			if preH != vm.RegDef.H.Value {
				flushTerminal = true
			}

			if preHirq != vm.RegIrq.H.Value {
				flushTerminal = true
			}

			if preIrqEn != vm.IrqEn {
				flushTerminal = true
			}

			if preInIrq != vm.InIrq {
				flushTerminal = true
			}

			// IPS counter
			stepCounter++
			if lastSecond != time.Now().Second() {
				lastSecond = time.Now().Second()
				currentSpeed = stepCounter
				stepCounter = 0
				flushTerminal = true
			}
		}

		writeVMState(vm, height, currentSpeed)

		// Sync terminal output
		if flushTerminal || vgaChanged || vm.Halted {
			termbox.Flush()
			vgaChanged = false
		}

		if len(closeChan) > 0 {
			return
		}
	}
}

func writeVMState(vm *VM, height, speed int) {
	state := ""
	if vm.Halted {
		state += " Halted "
	} else {
		state += " <Running>"
	}

	state += fmt.Sprintf(" | H <h%04X>", vm.RegDef.H.Value)
	state += fmt.Sprintf(" | H_irq <h%04X>", vm.RegIrq.H.Value)

	state += fmt.Sprintf(" | irq_en <%5t>", vm.IrqEn)
	state += fmt.Sprintf(" | in_irq <%5t>", vm.InIrq)
	state += fmt.Sprintf(" | #q_irq <%3d>", len(vm.IrqQueue))

	if speed > 1000000 {
		state += fmt.Sprintf(" | <%6.3f> MIPS", float64(speed)/1000000)
	} else if speed > 1000 {
		state += fmt.Sprintf(" | <%6.3f> kIPS", float64(speed)/1000)
	} else {
		state += fmt.Sprintf(" | <%7d>  IPS", speed)
	}

	state += fmt.Sprintf(" | con <%d>x<%d>", vm.VgaWidth, vm.VgaHeight)

	fg := termbox.ColorWhite
	runes := []rune(state)
	x := 0
	for i := 0; i < len(state); i++ {
		if runes[i] == '<' {
			fg = termbox.ColorCyan
		} else if runes[i] == '>' {
			fg = termbox.ColorWhite
		} else {
			termbox.SetCell(x+1, height-2, runes[i], fg, termbox.ColorBlack)
			x++
		}
	}
}
