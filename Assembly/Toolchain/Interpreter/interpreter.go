package interpreter

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Interpret runs a .mb binary on a virtual machine
func Interpret(file, config string, gui bool) {
	log.Println("Interpreting " + file + "...")
	if gui {
		log.Println("TODO")
	} else {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalln("ERROR: An error occured reading the input file.")
		}

		// Parse data into instruction-bounded array
		data16 := make([]uint16, len(data)/2)
		for i := 0; i < len(data16); i++ {
			data16[i] = uint16(data[i*2])<<8 | uint16(data[i*2+1])
		}

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
