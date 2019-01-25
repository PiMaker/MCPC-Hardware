package interpreter

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/nsf/termbox-go"
)

var (
	keycodeLookupRunes = map[rune]uint32{
		'a': 0x1c,
		'b': 0x32,
		'c': 0x21,
		'd': 0x23,
		'e': 0x24,
		'f': 0x2b,
		'g': 0x34,
		'h': 0x33,
		'i': 0x43,
		'j': 0x3b,
		'k': 0x42,
		'l': 0x4b,
		'm': 0x3a,
		'n': 0x31,
		'o': 0x44,
		'p': 0x4d,
		'q': 0x15,
		'r': 0x2d,
		's': 0x1b,
		't': 0x2c,
		'u': 0x3c,
		'v': 0x2a,
		'w': 0x1d,
		'x': 0x22,
		'z': 0x35,
		'y': 0x1a,

		'0': 0x45,
		'1': 0x16,
		'2': 0x1e,
		'3': 0x26,
		'4': 0x25,
		'5': 0x2e,
		'6': 0x36,
		'7': 0x3d,
		'8': 0x3e,
		'9': 0x46,

		'!': 0x16,
		'"': 0x1e,
		'§': 0x26,
		'$': 0x25,
		'%': 0x2e,
		'&': 0x36,
		'/': 0x3d,
		'(': 0x3e,
		')': 0x46,
		'=': 0x45,

		'-':  0x4a,
		'_':  0x4a,
		'ß':  0x4e,
		'?':  0x4e,
		',':  0x41,
		';':  0x41,
		'.':  0x49,
		':':  0x49,
		'ä':  0x52,
		'ö':  0x4c,
		'ü':  0x54,
		'+':  0x5b,
		'*':  0x5b,
		'#':  0x5d,
		'\'': 0x5d,
		'´':  0x55,
		'`':  0x55,
	}

	keycodeLookupRunesSpecialShifted = map[rune]uint32{
		'!': 0x16,
		'"': 0x1e,
		'§': 0x26,
		'$': 0x25,
		'%': 0x2e,
		'&': 0x36,
		'/': 0x3d,
		'(': 0x3e,
		')': 0x46,
		'=': 0x45,

		'_':  0x4e,
		'?':  0x55,
		';':  0x41,
		':':  0x49,
		'*':  0x5b,
		'\'': 0x5d,
		'`':  0x55,
	}

	keycodeLookupTermbox = map[termbox.Key]uint32{
		termbox.KeyBackspace:  0x66,
		termbox.KeyBackspace2: 0x66,
		termbox.KeyEnter:      0x5A,
		termbox.KeyEsc:        0x76,
		termbox.KeyTab:        0x0D,
		termbox.KeySpace:      0x29,
	}

	keycodeBreak     = uint32(0x00F00000)
	keycodeLSHFT     = uint32(0x00120000)
	keyboardIrqNum   = uint32(0xA)
	invalidKeyIrqNum = uint32(0xB)
)

// VMRun executes the given file in a virtual MCPC
func VMRun(file, traceFile string) {

	log.Println("Starting VM...")

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
			if x == 0 || x == width-1 || (x == 121 && y < 67 && y > 0 && y < height-3) {
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
	cpuIrqChan := make(chan uint32, 256)
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
						cpuIrqChan <- keyCode | keyboardIrqNum

						// Send BREAK
						cpuIrqChan <- keycodeBreak | keyboardIrqNum
						cpuIrqChan <- keyCode | keyboardIrqNum
					} else {
						// Invalid key interrupt
						cpuIrqChan <- invalidKeyIrqNum
					}
				} else {
					// Character key
					letter := event.Ch
					sentShift := false

					uppercase := unicode.IsUpper(letter)
					if uppercase {
						// Uppercase letter, send shift first
						cpuIrqChan <- keycodeLSHFT | keyboardIrqNum
						sentShift = true
						letter = unicode.ToLower(letter)
					}

					// Check for and perform special SHIFT handling
					if _, ok := keycodeLookupRunesSpecialShifted[letter]; ok {
						cpuIrqChan <- keycodeLSHFT | keyboardIrqNum
						sentShift = true
					}

					keyCode, ok := keycodeLookupRunes[letter]
					keyCode = keyCode << 16
					if ok {
						// Send MAKE
						cpuIrqChan <- keyCode | keyboardIrqNum

						// Send BREAK
						cpuIrqChan <- keycodeBreak | keyboardIrqNum
						cpuIrqChan <- keyCode | keyboardIrqNum
					} else {
						// Invalid key interrupt
						cpuIrqChan <- invalidKeyIrqNum
					}

					// Send SHFT BREAK in case of uppercase or special shift handling
					if sentShift {
						cpuIrqChan <- keycodeBreak | keyboardIrqNum
						cpuIrqChan <- keycodeLSHFT | keyboardIrqNum
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

	// Trace-handler
	var f *os.File
	if traceFile != "" {
		f, err = os.Create(traceFile)
		if err != nil {
			log.Fatalln("ERROR: " + err.Error())
		}

		f.WriteString(time.Now().Format(time.RFC3339) + " CPU tracing started. VM loaded file: " + file + "\n")
		vm.TraceCallback = func(msg string, step int64) {
			f.WriteString(time.Now().Format(time.RFC3339) + " [" + strconv.FormatInt(step, 10) + "] " + msg + "\n")
		}

		defer f.Close()
	}

	// Attach VGA update handler
	vgaChanged := false
	vm.VgaChangeCallback = func(addr, x, y, old, new uint16) {
		r := rune(new & 0x00FF)

		// Treat special characters like errors
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
		for len(cpuIrqChan) > 0 {
			vm.InjectIRQ(<-cpuIrqChan)
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
				if f != nil {
					f.Close()
				}
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
		state += fmt.Sprintf(" | <%7.3f> MIPS", float64(speed)/1000000)
	} else if speed > 1000 {
		state += fmt.Sprintf(" | <%7.3f> kIPS", float64(speed)/1000)
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
