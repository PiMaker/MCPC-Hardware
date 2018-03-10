package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./Compiler"
	"./Interpreter"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `MCPC Toolchain (Compiler/Linker/VM).

Usage:
  mcpc compile <file> <output> [--library=<library>...] [--offset=<offset>] [--enable-offset-jump] [--ascii]
  mcpc link <main> <output> --app=<include>... [--library=<library>...] [--ascii]
  mcpc interpret <file> [--config=<config>]
  mcpc debug <file> [--config=<config>]
  mcpc -h | --help
  mcpc --version

Options:
  compile                 Compiles an assembly file to binary.
  link                    Links multiple different assembly files together with an operating system (main) to create a static OS with program loading capabilities. Refer to documentation for program jump table format.
  interpret               Runs an MCPC virtual machine and executes a specified binary file. Use the --config flag to specify bus devices.
  debug                   Uses the MCPC interpreter to run the specified binary file and shows a TUI interface for debugging purposes.
  --library=<library>     Includes a library, specified in HJSON format, which allows higher-level instructions to be compiled down.
  --offset=<offset>       Specifies an offset that will be applied to the binary file [default: 0].
  --enable-offset-jump    If enabled, a 'jmp' instruction will be inserted at the beginning, jumping to the offset position. If the offset is smaller than 3, this flag will be ignored.
  --ascii                 Outputs the ascii binary format for use with the Digital circuit simulator.
  -h --help               Show this screen.
  --version               Show version.`

	// Parse command line arguments
	args, _ := docopt.ParseArgs(usage, os.Args[1:], "MineCraft PC Assembly Toolchain - Version 0.1")

	// Choose function to call based on arguments
	if argBool(args, "compile") {

		// Compile
		offset := argInt(args, "--offset")
		output := argString(args, "<output>")
		assembly := compiler.Compile(argString(args, "<file>"), offset, argStrings(args, "--library"), argBool(args, "--enable-offset-jump"))

		if argBool(args, "--ascii") {
			log.Println("Converting to ASCII format...")
			assembly = toASCIIFormat(assembly)
		}

		ioutil.WriteFile(output, assembly, 0666)

	} else if argBool(args, "link") {

		// Link
		log.Println("Linking OS () into ")

	} else if argBool(args, "interpret") || argBool(args, "debug") {

		// Interpret/Debug
		interpreter.Interpret(argString(args, "<file>"), "", argBool(args, "debug"))

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
		panic("Invalid argument \"" + key + "\"")
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
