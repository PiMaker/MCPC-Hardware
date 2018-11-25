package interpreter

import (
	"errors"
	"strconv"

	"github.com/tarm/serial"
)

// OP-Codes
const DEBUGGER_OPCODE_GET uint8 = 0x1
const DEBUGGER_OPCODE_SET uint8 = 0x2
const DEBUGGER_OPCODE_HI uint8 = 0x4
const DEBUGGER_OPCODE_LO uint8 = 0x8
const DEBUGGER_OPCODE_STEP uint8 = 0xC
const DEBUGGER_OPCODE_DUMP_ROM uint8 = 0xE
const DEBUGGER_OPCODE_DUMP_REGS uint8 = 0xA

// Device represents a physical MCPC device attached via a serial port
type Device struct {
	port *serial.Port
}

func establishSerialConnection(path string) (*Device, error) {
	port, err := serial.OpenPort(&serial.Config{
		Baud:     115200,
		Parity:   serial.ParityNone,
		Name:     path,
		StopBits: 1,
	})

	if err != nil {
		return nil, err
	}

	return &Device{
		port: port,
	}, nil
}

func (device *Device) closeConnection() {
	// Ignore error on close
	device.port.Close()
}

func (device *Device) waitForResponse() error {
	// Wait for feedback, discard received bytes though
	buf := make([]byte, 1)
	_, err := device.port.Read(buf)
	if err != nil {
		return err
	}
	return nil
}

func (device *Device) triggerROMDump() error {
	n, err := device.port.Write([]byte{DEBUGGER_OPCODE_DUMP_ROM})

	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("No bytes have been written")
	}

	return nil
}

func (device *Device) getRegister(reg uint8) (uint8, error) {
	n, err := device.port.Write([]byte{(reg << 4) | DEBUGGER_OPCODE_GET})
	if err != nil {
		return 0, err
	}

	if n <= 0 {
		return 0, errors.New("No bytes have been written")
	}

	buf := make([]byte, 1)
	n, err = device.port.Read(buf)

	if err != nil {
		return 0, err
	}

	if n <= 0 {
		return 0, errors.New("No data has been read")
	}

	return uint8(buf[0]), nil
}

func (device *Device) setRegister(reg, data uint8) error {
	n, err := device.port.Write([]byte{(reg << 4) | DEBUGGER_OPCODE_SET})
	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("No bytes have been written")
	}

	err = device.waitForResponse()
	if err != nil {
		return err
	}

	n, err = device.port.Write([]byte{(data & 0xF0) | DEBUGGER_OPCODE_HI})

	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("No bytes have been written")
	}

	err = device.waitForResponse()
	if err != nil {
		return err
	}

	n, err = device.port.Write([]byte{((data & 0x0F) << 4) | DEBUGGER_OPCODE_LO})

	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("No bytes have been written")
	}

	err = device.waitForResponse()
	if err != nil {
		return err
	}

	return nil
}

func (device *Device) step() error {
	n, err := device.port.Write([]byte{DEBUGGER_OPCODE_STEP})
	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New("No bytes have been written")
	}

	err = device.waitForResponse()
	if err != nil {
		return err
	}

	return nil
}

// NOTE: Disables instruction overwrite!
func (device *Device) getMCPCReg(reg uint8) (uint16, error) {
	err := device.setRegister(0, (1<<0)|(1<<1))
	if err != nil {
		return 0, err
	}
	err = device.setRegister(1, 0x0F&reg)
	if err != nil {
		return 0, err
	}
	low, err := device.getRegister(8)
	if err != nil {
		return 0, err
	}
	high, err := device.getRegister(9)
	if err != nil {
		return 0, err
	}
	regVal := uint16(low) | (uint16(high) << 8)
	err = device.setRegister(0, 1)
	if err != nil {
		return 0, err
	}
	return regVal, nil
}

func (device *Device) getMCPCRegDump() ([]uint16, error) {
	n, err := device.port.Write([]byte{DEBUGGER_OPCODE_DUMP_REGS})
	if err != nil {
		return nil, err
	}

	if n <= 0 {
		return nil, errors.New("No bytes have been written")
	}

	dump := make([]uint16, 16)
	rawDump := make([]byte, 32)

	n, err = device.port.Read(rawDump)
	if err != nil {
		return nil, err
	}

	if n != len(rawDump) {
		return nil, errors.New("Regdump: Wrong amount of bytes read: " + strconv.Itoa(n))
	}

	for i := 0; i < 16; i++ {
		dump[i] = uint16(rawDump[i*2]) | (uint16(rawDump[i*2+1]) << 8)
	}

	return dump, nil
}
