package interpreter

import (
	"errors"
	"fmt"
)

const (
	regFrom uint16 = 0x00F0
	regData uint16 = 0x00F0
	regTo   uint16 = 0x0F00
	regIf   uint16 = 0xF000
	regOp   uint16 = 0xF000
)

// VM represents an MCPC virtual machine state
type VM struct {
	Registers *Registers
	SRAM      []uint16
	EEPROM    []uint16
	Halted    bool
}

// Registers includes all registers of an MCPC instance
type Registers struct {
	A, B, C, D, E, F, G, H, SCR1, SCR2, SP, PC, Zero, One, NegOne, BUS *Register
}

// Register represents a single register of an MCPC instance
type Register struct {
	Address   byte
	Value     uint16
	Writeable bool
}

// MaxSRAMValue is the maximum address the SRAM can be accessed at
const MaxSRAMValue uint16 = 0x7FFF

var sramWriteWaiting = false
var sramWriteAddress uint16

// NewVM creates a new MCPC virtual machine instance
func NewVM(program []uint16) *VM {

	sram := make([]uint16, MaxSRAMValue+1)

	regs := &Registers{
		A:      &Register{Value: 0, Address: 0x0, Writeable: true},
		B:      &Register{Value: 0, Address: 0x1, Writeable: true},
		C:      &Register{Value: 0, Address: 0x2, Writeable: true},
		D:      &Register{Value: 0, Address: 0x3, Writeable: true},
		E:      &Register{Value: 0, Address: 0x4, Writeable: true},
		F:      &Register{Value: 0, Address: 0x5, Writeable: true},
		G:      &Register{Value: 0, Address: 0x6, Writeable: true},
		H:      &Register{Value: 0, Address: 0x7, Writeable: true},
		SCR1:   &Register{Value: 0, Address: 0x8, Writeable: true},
		SCR2:   &Register{Value: 0, Address: 0x9, Writeable: true},
		SP:     &Register{Value: 0, Address: 0xA, Writeable: true},
		PC:     &Register{Value: 0, Address: 0xB, Writeable: true},
		Zero:   &Register{Value: 0, Address: 0xC, Writeable: false},
		One:    &Register{Value: 1, Address: 0xD, Writeable: false},
		NegOne: &Register{Value: 0xFFFF, Address: 0xE, Writeable: false},
		BUS:    &Register{Value: 0, Address: 0xF, Writeable: false},
	}

	vm := VM{
		EEPROM:    program,
		SRAM:      sram,
		Registers: regs,
		Halted:    false,
	}

	return &vm
}

// Step executes a single instruction step of this MCPC virtual machine instance; Returns true if a debug break instruction has been hit
func (vm *VM) Step() (bool, string, error) {
	if vm.Halted {
		return false, "", nil
	}

	brk := false
	termout := ""
	var err error

	if int(vm.Registers.PC.Value) >= len(vm.EEPROM) {
		return false, "", errors.New("Invalid EEPROM address, PC out of range")
	}

	ins := vm.EEPROM[vm.Registers.PC.Value]
	instruction := ins & 0x000F

	switch instruction {
	case 0x0:
		vm.Halted = true
	case 0x1:
		reg := getReg(vm, ins, regTo)
		if reg.Writeable {
			reg.Value = getReg(vm, ins, regFrom).Value

			if reg == vm.Registers.PC {
				vm.Registers.PC.Value--
			}
		} else {
			err = fmt.Errorf("Write to non-writable register %X", reg.Address)
		}
	case 0x2:
		if getReg(vm, ins, regIf).Value != 0 {
			reg := getReg(vm, ins, regTo)
			if reg.Writeable {
				reg.Value = getReg(vm, ins, regFrom).Value

				if reg == vm.Registers.PC {
					vm.Registers.PC.Value--
				}
			} else {
				err = fmt.Errorf("Write to non-writable register %X", reg.Address)
			}
		}
	case 0x3:
		if getReg(vm, ins, regIf).Value == 0 {
			reg := getReg(vm, ins, regTo)
			if reg.Writeable {
				reg.Value = getReg(vm, ins, regFrom).Value

				if reg == vm.Registers.PC {
					vm.Registers.PC.Value--
				}
			} else {
				err = fmt.Errorf("Write to non-writable register %X", reg.Address)
			}
		}

	case 0x4:
		// BUS is deprecated, use as break in VM
		brk = true

	case 0x5:
		addrReg := getReg(vm, ins, regFrom)
		writeToReg := getReg(vm, ins, regTo)

		if writeToReg.Writeable {
			if (addrReg.Value & 0x8000) == 0 {
				writeToReg.Value = vm.SRAM[addrReg.Value]
			} else if addrReg.Value == 0x8000 {
				writeToReg.Value = 0x8001
			} else if addrReg.Value == 0x8065 {
				writeToReg.Value = 0xE000
			} else if addrReg.Value >= 0xD000 && addrReg.Value < 0xD800 {
				writeToReg.Value = vm.EEPROM[addrReg.Value-0xD000]
			} else {
				// other CFGs return 0 (not implemented)
				writeToReg.Value = 0
			}
		} else {
			err = fmt.Errorf("Write to non-writable register %X", writeToReg.Address)
		}

	case 0x6:
		vm.Registers.PC.Value++
		reg := getReg(vm, ins, regTo)
		if reg == vm.Registers.PC {
			err = errors.New("SET was called on PC, this is not allowed in VM")
		} else {
			if reg.Writeable {
				reg.Value = vm.EEPROM[vm.Registers.PC.Value]
			} else {
				err = fmt.Errorf("Write to non-writable register %X", reg.Address)
			}
		}

	case 0x7:
		addrReg := getReg(vm, ins, regFrom)
		dataReg := getReg(vm, ins, regIf)

		if (addrReg.Value & 0x8000) == 0 {
			vm.SRAM[addrReg.Value] = dataReg.Value
		} else {
			// ignore CFG writes for now
		}

	case 0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF:
		registerTo := getReg(vm, ins, regTo)
		if registerTo.Writeable {
			// ALU instruction decoding
			switch instruction {
			case 0x8:
				registerTo.Value = getReg(vm, ins, regFrom).Value & getReg(vm, ins, regOp).Value
			case 0x9:
				registerTo.Value = getReg(vm, ins, regFrom).Value | getReg(vm, ins, regOp).Value
			case 0xA:
				registerTo.Value = getReg(vm, ins, regFrom).Value ^ getReg(vm, ins, regOp).Value
			case 0xB:
				registerTo.Value = getReg(vm, ins, regFrom).Value + getReg(vm, ins, regOp).Value
			case 0xC:
				shft := (ins & regOp) >> 12
				if shft&0x8 == 0 {
					registerTo.Value = getReg(vm, ins, regFrom).Value >> shft
				} else {
					registerTo.Value = getReg(vm, ins, regFrom).Value << (shft & 0x7)
				}
			case 0xD:
				registerTo.Value = getReg(vm, ins, regFrom).Value * getReg(vm, ins, regOp).Value
			case 0xE:
				var val uint16
				if getReg(vm, ins, regFrom).Value > getReg(vm, ins, regOp).Value {
					val = 0xFFFF
				}
				registerTo.Value = val
			case 0xF:
				var val uint16
				if getReg(vm, ins, regFrom).Value == getReg(vm, ins, regOp).Value {
					val = 0xFFFF
				}
				registerTo.Value = val
			}
		} else {
			err = fmt.Errorf("Write to non-writable register %X (by ALU)", registerTo.Address)
		}
	}

	// Increase program counter by one
	if !vm.Halted {
		vm.Registers.PC.Value++
	}

	return brk, termout, err
}

func getReg(vm *VM, ins uint16, reg uint16) *Register {
	addr := ins & reg
	addr >>= getFirstSet(reg)
	switch byte(addr) {
	case vm.Registers.A.Address:
		return vm.Registers.A
	case vm.Registers.B.Address:
		return vm.Registers.B
	case vm.Registers.C.Address:
		return vm.Registers.C
	case vm.Registers.D.Address:
		return vm.Registers.D
	case vm.Registers.E.Address:
		return vm.Registers.E
	case vm.Registers.F.Address:
		return vm.Registers.F
	case vm.Registers.G.Address:
		return vm.Registers.G
	case vm.Registers.H.Address:
		return vm.Registers.H
	case vm.Registers.SCR1.Address:
		return vm.Registers.SCR1
	case vm.Registers.SCR2.Address:
		return vm.Registers.SCR2
	case vm.Registers.SP.Address:
		return vm.Registers.SP
	case vm.Registers.PC.Address:
		return vm.Registers.PC
	case vm.Registers.Zero.Address:
		return vm.Registers.Zero
	case vm.Registers.One.Address:
		return vm.Registers.One
	case vm.Registers.NegOne.Address:
		return vm.Registers.NegOne
	case vm.Registers.BUS.Address:
		return vm.Registers.BUS
	default:
		return vm.Registers.A
	}
}

func getFirstSet(val uint16) byte {
	var ret byte
	for val&0x1 == 0 {
		val >>= 1
		ret++
	}
	return ret
}

func (vm *VM) registerFromNumber(regNum uint16) *Register {
	return getReg(vm, regNum, 0xFFFF)
}

func (vm *VM) compareRegistersWithDevice(dev *Device) (different bool, differences []string, err error) {
	diff := make([]string, 0)

	dump, err := dev.getMCPCRegDump()
	if err != nil {
		return false, nil, err
	}

	for reg := uint16(0); reg < 16; reg++ {
		devReg := dump[reg]
		vmReg := vm.registerFromNumber(reg)

		if vmReg.Value != devReg {
			diff = append(diff, fmt.Sprintf("Register %d: VM=0x%04x : Device=0x%04x", reg, vmReg.Value, devReg))
		}
	}

	return len(diff) > 0, diff, nil
}
