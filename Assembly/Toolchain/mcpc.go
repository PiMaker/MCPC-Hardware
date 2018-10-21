package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./Assembler"
	"./Interpreter"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `MCPC Toolchain (Compiler/Assembler/Linker/VM).

Usage:
  mcpc assemble <file> <output> [--library=<library>...] [--debug-symbols] [--offset=<offset>] [--enable-offset-jump] [--ascii] [--hex] [--length=<length>]
  mcpc interpret <file> [--max-steps=<max-steps>] [--config=<config>]
  mcpc debug <file> [--config=<config>]
  mcpc -h | --help
  mcpc --version

Options:
  assemble                Assembles an assembler file to assembly.
  interpret               Runs an MCPC virtual machine and executes a specified binary file. Use the --config flag to specify bus devices.
  debug                   Uses the MCPC interpreter to run the specified binary file and shows a TUI interface for debugging purposes.
  --library=<library>     Includes a library, specified in HJSON format, which allows higher-level instructions to be compiled down.
  --debug-symbols         Writes a symbol file to use with the MCPC debugger next to the output file (will overwrite existing symbol files!)
  --offset=<offset>       Specifies an offset that will be applied to the binary file [default: 0].
  --enable-offset-jump    If enabled, a 'jmp' instruction will be inserted at the beginning, jumping to the offset position. If the offset is smaller than 3, this flag will be ignored.
  --ascii                 Outputs the ascii binary format for use with the Digital circuit simulator.
  --hex                   Outputs raw binary in Verilog HEX format.
  --length=<length>       Length of hex output in bytes (one instruction is 2 bytes!) [default: 4096].
  --max-steps=<max-steps> Sets a maximum amount of steps for interpreting a binary file [default: 100000].
  -h --help               Show this screen.
  --version               Show version.`

	// Parse command line arguments
	args, _ := docopt.ParseArgs(usage, os.Args[1:], "MCPC Assembly Toolchain - Version 0.1")

	// Choose function to call based on arguments
	if argBool(args, "assemble") {

		if argBool(args, "--ascii") && argBool(args, "--hex") {
			panic("Can only specify one alternate output format (ASCII/HEX)")
		}

		// Compile
		offset := argInt(args, "--offset")
		output := argString(args, "<output>")
		assembly, debugSymbols := assembler.Compile(argString(args, "<file>"), offset, argStrings(args, "--library"), argBool(args, "--enable-offset-jump"))

		if argBool(args, "--ascii") {
			log.Println("Converting to ASCII format...")
			assembly = toASCIIFormat(assembly)
		} else if argBool(args, "--hex") {
			assembly = toHEXFormat(assembly, argInt(args, "--length"))
		}

		ioutil.WriteFile(output, assembly, 0664)

		if argBool(args, "--debug-symbols") {
			symbolFile := output + ".msym"
			ioutil.WriteFile(symbolFile, debugSymbols, 0664)
		}

	} else if argBool(args, "interpret") || argBool(args, "debug") {

		// Interpret/Debug
		interpreter.Interpret(argString(args, "<file>"), "", argBool(args, "debug"), argInt(args, "--max-steps"))

	} else {
		log.Println("Invalid command, use -h for help")
	}
}

func argString(args docopt.Opts, key string) string {
	v, err := args.String(key)
	if err != nil {
		panic("Invalid argument \"" + key + "\"")
	}

	return v
}

func argStrings(args docopt.Opts, key string) []string {
	v, err := args[key].([]string)
	if !err {
		return make([]string, 0)
	}

	return v
}

func argBool(args docopt.Opts, key string) bool {
	v, err := args.Bool(key)
	if err != nil {
		panic("Invalid argument \"" + key + "\"")
	}

	return v
}

func argInt(args docopt.Opts, key string) int {
	v, err := args.Int(key)
	if err != nil {
		// No panic here, just trust me on this
		return -1
	}

	return v
}

func toASCIIFormat(data []byte) []byte {
	header := []byte("v2.0 raw\n")
	retval := make([]byte, len(header)+len(data)*3)

	marker := len(header)
	copy(retval, header)

	for i := 0; i < len(data); i++ {
		val := []byte(fmt.Sprintf("%x\n", data[i]))
		retval[marker] = val[0]
		retval[marker+1] = val[1]
		if len(val) > 2 {
			retval[marker+2] = val[2]
		}
		marker += len(val)
	}

	return retval[:marker]
}

func toHEXFormat(data []byte, length int) []byte {
	var retval bytes.Buffer

	log.Printf("Converting to Verilog hex, padding to: %d\n", length)

	buf := make([]byte, length)
	copy(buf, data)

	// Theoretically shouldn't happen, but better safe than sorry
	if len(buf)%2 != 0 {
		buf = append(buf, 0)
	}

	for i := 0; i < len(buf); i += 2 {
		retval.WriteString(fmt.Sprintf("%02x%02x\n", buf[i], buf[i+1]))
	}

	return retval.Bytes()
}
