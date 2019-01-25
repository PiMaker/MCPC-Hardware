package interpreter

import (
	"errors"
	"fmt"
	"unsafe"
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
	RegDef      *Registers
	RegIrq      *Registers
	SRAM        []uint16
	EEPROM      []uint16
	Halted      bool
	SRAMPageDef uint16
	SRAMPageIrq uint16
	VgaWidth    uint16
	VgaHeight   uint16
	VgaBuffer   []uint16

	IrqEn      bool
	InIrq      bool
	IrqHandler uint16
	IrqQueue   chan uint32
	irqDataBuf uint32

	// addr relative to VgaBuffer (0 is x=y=0)
	VgaChangeCallback func(addr, x, y, old, new uint16)

	TraceCallback func(msg string, step int64)
	StepCounter   int64
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

// SRAMPageCount determines how many memory pages are available
const SRAMPageCount uint = 16

// SRAMPageBits is the bit count needed to represent SRAMPageCount
const SRAMPageBits uint = 4

// SRAMPageMask is a mask that can be applied to a 16 bit unsigned integer to make it safe to use as an SRAM page address
const SRAMPageMask uint16 = 0x000F

var sramWriteWaiting = false
var sramWriteAddress uint16

// Register arbitration (DEF/IRQ)
func (vm *VM) Registers() *Registers {
	if vm.InIrq {
		return vm.RegIrq
	}

	return vm.RegDef
}

// SRAMPage arbitration (DEF/IRQ)
func (vm *VM) SRAMPage() uint16 {
	if vm.InIrq {
		return vm.SRAMPageIrq
	}

	return vm.SRAMPageDef
}

// NewVM creates a new MCPC virtual machine instance
func NewVM(program []uint16, vgaWidth, vgaHeight uint16) *VM {

	// Clamp to max VGA size to avoid overflow
	if vgaWidth > 120 {
		vgaWidth = 120
	}
	if vgaHeight > 65 {
		vgaHeight = 65
	}

	sram := make([]uint16, 0x1FFFFFF)

	vm := VM{
		EEPROM:      program,
		SRAM:        sram,
		RegDef:      generateEmptyRegisterSet(0),
		RegIrq:      generateEmptyRegisterSet(0),
		Halted:      false,
		SRAMPageDef: 0,
		SRAMPageIrq: 0,
		VgaWidth:    vgaWidth,
		VgaHeight:   vgaHeight,
		VgaBuffer:   make([]uint16, vgaWidth*vgaHeight),

		InIrq:      false,
		IrqEn:      false,
		IrqHandler: 0,
		IrqQueue:   make(chan uint32, 256),

		StepCounter: -1,
	}

	return &vm
}

// Step executes a single instruction step of this MCPC virtual machine instance; Returns true if a debug break instruction has been hit
func (vm *VM) Step() (bool, error) {
	if vm.Halted {
		return false, nil
	}

	vm.StepCounter++

	brk := false
	var err error

	// Check for IRQ enter
	if vm.IrqEn && !vm.InIrq && len(vm.IrqQueue) > 0 {
		vm.RegIrq = generateEmptyRegisterSet(vm.IrqHandler)
		vm.InIrq = true
		vm.irqDataBuf = <-vm.IrqQueue
		vm.SRAMPageIrq = 0

		vm.t("IRQ enter, buf=h%08X, irq_reg_pc=h%04X (irq_handler=h%04X), #irq_q=%d", vm.irqDataBuf, vm.RegIrq.PC.Value, vm.IrqHandler, len(vm.IrqQueue))
	}

	if int(vm.Registers().PC.Value) >= len(vm.EEPROM) {
		vm.t("Invalid EEPROM address: pc=h%04X, len(EEPROM)=h%X", vm.Registers().PC.Value, len(vm.EEPROM))
		return false, errors.New("Invalid EEPROM address, PC out of range")
	}

	ins := vm.EEPROM[vm.Registers().PC.Value]
	instruction := ins & 0x000F

	vm.t("Ins: EEPROM[h%04X]=h%04X (irq=%t)", vm.Registers().PC.Value, ins, vm.InIrq)
	vm.tReg()

	switch instruction {
	case 0x0:
		vm.t("HALT, state now halted")
		vm.Halted = true
	case 0x1:
		vm.t("MOV")
		reg := GetReg(vm, ins, regTo)
		if reg.Writeable {
			reg.Value = GetReg(vm, ins, regFrom).Value

			if reg == vm.Registers().PC {
				vm.Registers().PC.Value--
			}
		} else {
			err = fmt.Errorf("Write to non-writable register %X", reg.Address)
		}
	case 0x2:
		if GetReg(vm, ins, regIf).Value != 0 {
			vm.t("MOVNZ (applied)")
			reg := GetReg(vm, ins, regTo)
			if reg.Writeable {
				reg.Value = GetReg(vm, ins, regFrom).Value

				if reg == vm.Registers().PC {
					vm.Registers().PC.Value--
				}
			} else {
				err = fmt.Errorf("Write to non-writable register %X", reg.Address)
			}
		} else {
			vm.t("MOVNZ (skipped)")
		}
	case 0x3:
		if GetReg(vm, ins, regIf).Value == 0 {
			vm.t("MOVEZ (applied)")
			reg := GetReg(vm, ins, regTo)
			if reg.Writeable {
				reg.Value = GetReg(vm, ins, regFrom).Value

				if reg == vm.Registers().PC {
					vm.Registers().PC.Value--
				}
			} else {
				err = fmt.Errorf("Write to non-writable register %X", reg.Address)
			}
		} else {
			vm.t("MOVEZ (skipped)")
		}

	case 0x4:
		// BUS is deprecated, use as break in VM
		brk = true
		vm.t("BUS (deprecated, brk)")

	case 0x5:
		addrReg := GetReg(vm, ins, regFrom)
		writeToReg := GetReg(vm, ins, regTo)

		vm.t("MEMR:")

		if writeToReg.Writeable {
			if (addrReg.Value & 0x8000) == 0 {
				// Direct
				writeToReg.Value = vm.SRAM[uint(vm.SRAMPage()&SRAMPageMask)<<16|uint(addrReg.Value)]
				vm.t("Direct, SRAM[h%07X] == h%04X", uint(vm.SRAMPage()&SRAMPageMask)<<16|uint(addrReg.Value), writeToReg.Value)
			} else if addrReg.Value == 0x8000 {
				vm.t("CFG, Version (0x8001)")
				writeToReg.Value = 0x8001
			} else if addrReg.Value >= 0xD000 && addrReg.Value < 0xD800 {
				vm.t("CFG, Access to EEPROM (len=h%X): [h%X-0xD000=h%X]", len(vm.EEPROM), addrReg.Value, addrReg.Value-0xD000)
				writeToReg.Value = vm.EEPROM[addrReg.Value-0xD000]
			} else if addrReg.Value == 0x8800 {
				writeToReg.Value = vm.SRAMPage()
				vm.t("CFG, SRAMPage = h%04X", writeToReg.Value)
			} else if addrReg.Value == 0xDFFD {
				// VGA
				writeToReg.Value = vm.VgaWidth
				vm.t("CFG, VgaWidth = h%04X", writeToReg.Value)
			} else if addrReg.Value == 0xDFFE {
				writeToReg.Value = vm.VgaHeight
				vm.t("CFG, VgaHeight = h%04X", writeToReg.Value)
			} else if addrReg.Value == 0xDFFF {
				writeToReg.Value = 0xE000 + vm.VgaWidth*vm.VgaHeight - 1
				vm.t("CFG, VgaMaxAddr = h%04X", writeToReg.Value)
			} else if addrReg.Value >= 0xE000 && addrReg.Value < uint16(0xE000+len(vm.VgaBuffer)) {
				// Read from ASCII buffer
				writeToReg.Value = vm.VgaBuffer[addrReg.Value-0xE000]
				vm.t("CFG, read from VGA buffer [abs=h%04X,rel=h%04X] == h%04X", addrReg.Value, addrReg.Value-0xE000, writeToReg.Value)
			} else if addrReg.Value == 0x9000 {
				// IRQs
				writeToReg.Value = vm.IrqHandler
				vm.t("CFG, IrqHandler = h%04X", writeToReg.Value)
			} else if addrReg.Value == 0x9001 {
				if vm.IrqEn {
					writeToReg.Value = 0xFFFF
				} else {
					writeToReg.Value = 0x0
				}
				vm.t("CFG, IRQ_EN = h%04X", writeToReg.Value)
			} else if addrReg.Value == 0x9002 {
				if vm.InIrq {
					writeToReg.Value = 0xFFFF
				} else {
					writeToReg.Value = 0x0
				}
				vm.t("CFG, IN_IRQ = h%04X", writeToReg.Value)
			} else if addrReg.Value == 0x9010 {
				if vm.InIrq {
					writeToReg.Value = uint16(vm.irqDataBuf)
					vm.t("CFG, IrqBuf_low = h%04X", writeToReg.Value)
				} else {
					writeToReg.Value = 0
					vm.t("CFG, IrqBuf_low = 0 (not in IRQ)")
				}
			} else if addrReg.Value == 0x9011 {
				if vm.InIrq {
					writeToReg.Value = uint16(vm.irqDataBuf >> 16)
					vm.t("CFG, IrqBuf_high = h%04X", writeToReg.Value)
				} else {
					writeToReg.Value = 0
					vm.t("CFG, IrqBuf_high = 0 (not in IRQ)")
				}
			} else {
				// other CFGs return 0 (not implemented)
				writeToReg.Value = 0
				vm.t("Unknown CFG: addr=h%04X", addrReg.Value)
			}
		} else {
			err = fmt.Errorf("Write to non-writable register %X", writeToReg.Address)
		}

	case 0x6:
		vm.Registers().PC.Value++
		reg := GetReg(vm, ins, regTo)
		if reg == vm.Registers().PC {
			err = errors.New("SET was called on PC, this is not allowed in VM")
		} else {
			if reg.Writeable {
				reg.Value = vm.EEPROM[vm.Registers().PC.Value]
				vm.t("SET reg%X <- EEPROM[h%04X]=h%04X", reg.Address, vm.Registers().PC.Value, reg.Value)
			} else {
				err = fmt.Errorf("Write to non-writable register 0x%X", reg.Address)
			}
		}

	case 0x7:
		addrReg := GetReg(vm, ins, regFrom)
		dataReg := GetReg(vm, ins, regIf)

		vm.t("MEMW:")

		if (addrReg.Value & 0x8000) == 0 {
			vm.SRAM[uint(vm.SRAMPage()&SRAMPageMask)<<16|uint(addrReg.Value)] = dataReg.Value
			vm.t("Direct, SRAM[h%07X] <- h%04X", uint(vm.SRAMPage()&SRAMPageMask)<<16|uint(addrReg.Value), dataReg.Value)
		} else if addrReg.Value >= 0xE000 && addrReg.Value < uint16(0xE000+len(vm.VgaBuffer)) {
			// VGA
			relAddr := addrReg.Value - 0xE000
			old := vm.VgaBuffer[relAddr]
			vm.VgaBuffer[relAddr] = dataReg.Value
			vm.t("CFG, VGA, write to addr=h%04X (rel=h%04X), x=%d y=%d rune=%v/('%s')", addrReg.Value, relAddr, relAddr%vm.VgaWidth, relAddr/vm.VgaWidth, rune(dataReg.Value), string(rune(dataReg.Value)))
			if vm.VgaChangeCallback != nil {
				vm.VgaChangeCallback(relAddr,
					relAddr%vm.VgaWidth,
					relAddr/vm.VgaWidth,
					old,
					dataReg.Value)
				vm.t("VGA change callback fired")
			}
		} else if addrReg.Value == 0x9000 { // IRQs
			vm.IrqHandler = dataReg.Value
			vm.t("CFG, IRQ handler <- h%04X", vm.IrqHandler)
		} else if addrReg.Value == 0x9001 {
			vm.IrqEn = dataReg.Value != 0
			vm.t("CFG, IRQ_EN <- %t", vm.IrqEn)

			if dataReg.Value == 0 {
				// Empty IRQ queue on IRQ disable
				vm.t("IRQ disable, %d queued IRQs dropped", len(vm.IrqQueue))
				for len(vm.IrqQueue) > 0 {
					<-vm.IrqQueue
				}
			}
		} else if addrReg.Value == 0x9002 {
			if dataReg.Value == 0 {
				vm.InIrq = false
				// Counter regular PC increment by decrementing first
				// Otherwise PC would be increased double after an IRQ
				vm.Registers().PC.Value--
				vm.t("CFG, IRQ exit")
			} else {
				vm.t("CFG, IRQ exit attempted, write value not 0 (was: h%04X)", dataReg.Value)
			}
		} else if addrReg.Value == 0xFFFF { // Debug BRK
			vm.t("CFG, BRK")
			brk = true
		} else if addrReg.Value == 0x8800 {
			if vm.InIrq {
				vm.SRAMPageIrq = dataReg.Value
			} else {
				vm.SRAMPageDef = dataReg.Value
			}
			vm.t("CFG, SRAMPage <- h%04X", dataReg.Value)
		} else {
			// CFG write, ignores unknown CFGs
			vm.t("Unknown CFG, addr=h%04X", addrReg.Value)
		}

	case 0x8, 0x9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF:
		registerTo := GetReg(vm, ins, regTo)
		if registerTo.Writeable {
			// ALU instruction decoding
			switch instruction {
			case 0x8:
				registerTo.Value = GetReg(vm, ins, regFrom).Value & GetReg(vm, ins, regOp).Value
				vm.t("ALU &")
			case 0x9:
				registerTo.Value = GetReg(vm, ins, regFrom).Value | GetReg(vm, ins, regOp).Value
				vm.t("ALU |")
			case 0xA:
				registerTo.Value = GetReg(vm, ins, regFrom).Value ^ GetReg(vm, ins, regOp).Value
				vm.t("ALU ^")
			case 0xB:
				registerTo.Value = GetReg(vm, ins, regFrom).Value + GetReg(vm, ins, regOp).Value
				vm.t("ALU +")
			case 0xC:
				// SHFT
				by := GetReg(vm, ins, regOp).Value
				if by&0xFF00 == 0 {
					registerTo.Value = GetReg(vm, ins, regFrom).Value >> (by & 0x00FF)
					vm.t("ALU >>")
				} else {
					registerTo.Value = GetReg(vm, ins, regFrom).Value << (by & 0x00FF)
					vm.t("ALU <<")
				}

			case 0xD:
				registerTo.Value = GetReg(vm, ins, regFrom).Value * GetReg(vm, ins, regOp).Value
				vm.t("ALU *")
			case 0xE:
				// Unsafe cast to int16 to preserve bit values (two's complement) for size checking
				var op1, op2 int16
				op1u := GetReg(vm, ins, regFrom).Value
				op2u := GetReg(vm, ins, regOp).Value
				op1 = *(*int16)(unsafe.Pointer(&op1u))
				op2 = *(*int16)(unsafe.Pointer(&op2u))

				var val uint16
				if op1 > op2 {
					val = 0xFFFF
				}
				registerTo.Value = val
				vm.t("ALU > (%t)", val == 0xFFFF)
			case 0xF:
				var val uint16
				if GetReg(vm, ins, regFrom).Value == GetReg(vm, ins, regOp).Value {
					val = 0xFFFF
				}
				registerTo.Value = val
				vm.t("ALU = (%t)", val == 0xFFFF)
			}
		} else {
			err = fmt.Errorf("Write to non-writable register 0x%X (by ALU)", registerTo.Address)
		}
	}

	// Increase program counter by one
	vm.Registers().PC.Value++

	vm.tReg()

	return brk, err
}

// InjectIRQ writes the given IRQ payload into the VM's IRQ FIFO
func (vm *VM) InjectIRQ(irqData uint32) {
	// Discard IRQ if queue full or IRQ disabled
	if vm.IrqEn && len(vm.IrqQueue) < cap(vm.IrqQueue) {
		vm.t("> Interrupt enqueued: h%08X", irqData)
		vm.IrqQueue <- irqData
	} else {
		vm.t("> Interrupt discarded: h%08X", irqData)
	}
}

// GetReg extracts details about a register from an instruction in the context of a VM
func GetReg(vm *VM, ins uint16, reg uint16) *Register {
	addr := ins & reg

	if reg == 0x00F0 {
		addr >>= 4
	} else if reg == 0x0F00 {
		addr >>= 8
	} else if reg == 0xF000 {
		addr >>= 12
	}

	switch byte(addr) {
	case vm.Registers().A.Address:
		return vm.Registers().A
	case vm.Registers().B.Address:
		return vm.Registers().B
	case vm.Registers().C.Address:
		return vm.Registers().C
	case vm.Registers().D.Address:
		return vm.Registers().D
	case vm.Registers().E.Address:
		return vm.Registers().E
	case vm.Registers().F.Address:
		return vm.Registers().F
	case vm.Registers().G.Address:
		return vm.Registers().G
	case vm.Registers().H.Address:
		return vm.Registers().H
	case vm.Registers().SCR1.Address:
		return vm.Registers().SCR1
	case vm.Registers().SCR2.Address:
		return vm.Registers().SCR2
	case vm.Registers().SP.Address:
		return vm.Registers().SP
	case vm.Registers().PC.Address:
		return vm.Registers().PC
	case vm.Registers().Zero.Address:
		return vm.Registers().Zero
	case vm.Registers().One.Address:
		return vm.Registers().One
	case vm.Registers().NegOne.Address:
		return vm.Registers().NegOne
	case vm.Registers().BUS.Address:
		return vm.Registers().BUS
	default:
		return vm.Registers().A
	}
}

func (vm *VM) t(f string, p ...interface{}) {
	if vm.TraceCallback != nil {
		vm.TraceCallback(fmt.Sprintf(f, p...), vm.StepCounter)
	}
}

func (vm *VM) tReg() {
	vm.t("REGDUMP; IRQ: %t; A=h%04X, B=h%04X, C=h%04X, D=h%04X, E=h%04X, F=h%04X, G=h%04X, H=h%04X, SCR1=h%04X, SCR2=h%04X, SP=h%04X, PC=h%04X, 0=h%04X, 1=h%04X, -1=h%04X, BUS=h%04X",
		vm.Registers() == vm.RegIrq,
		vm.Registers().A.Value,
		vm.Registers().B.Value,
		vm.Registers().C.Value,
		vm.Registers().D.Value,
		vm.Registers().E.Value,
		vm.Registers().F.Value,
		vm.Registers().G.Value,
		vm.Registers().H.Value,
		vm.Registers().SCR1.Value,
		vm.Registers().SCR2.Value,
		vm.Registers().SP.Value,
		vm.Registers().PC.Value,
		vm.Registers().Zero.Value,
		vm.Registers().One.Value,
		vm.Registers().NegOne.Value,
		vm.Registers().BUS.Value)
}

func (vm *VM) registerFromNumber(regNum uint16) *Register {
	return GetReg(vm, regNum, 0xFFFF)
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

func generateEmptyRegisterSet(pc uint16) *Registers {
	return &Registers{
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
		PC:     &Register{Value: pc, Address: 0xB, Writeable: true},
		Zero:   &Register{Value: 0, Address: 0xC, Writeable: false},
		One:    &Register{Value: 1, Address: 0xD, Writeable: false},
		NegOne: &Register{Value: 0xFFFF, Address: 0xE, Writeable: false},
		BUS:    &Register{Value: 0, Address: 0xF, Writeable: false},
	}
}
