package main

import (
	"fmt"
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
  mcpc -h | --help
  mcpc --version

Options:
  compile                 Compiles an assembly file to binary.
  link                    Links multiple different assembly files together with an operating system (main) to create a static OS with program loading capabilities. Refer to documentation for program jump table format.
  interpret               Runs an MCPC virtual machine and executes a specified binary file. Use the --config flag to specify bus devices.
  --library=<library>     Includes a library, specified in HJSON format, which allows higher-level instructions to be compiled down.
  --offset=<offset>       Specifies an offset that will be applied to the binary file [default: 0].
  --enable-offset-jump    If enabled, a 'jmp' instruction will be inserted at the beginning, jumping to the offset position. If the offset is smaller than 2, this flag will be ignored.
  --ascii                 Outputs the ascii binary format for use with the Digital circuit simulator.
  -h --help               Show this screen.
  --version               Show version.`

	// Parse command line arguments
	args, _ := docopt.ParseArgs(usage, os.Args[1:], "MineCraft PC Assembly Toolchain - Version 0.1")

	// Choose function to call based on arguments
	if argBool(args, "compile") {

		// Compile
		offset := argInt(args, "--offset")
		compiler.Compile(argString(args, "<file>"), argString(args, "<output>"), offset, make([]string, 0))

	} else if argBool(args, "link") {

		// Link
		fmt.Println("Linking OS () into ")

	} else if argBool(args, "interpret") {

		// Interpret
		interpreter.Interpret(argString(args, "<file>"), "")

	} else {
		fmt.Println("Invalid command, use -h for help")
	}
}

func argString(args docopt.Opts, key string) string {
	v, err := args.String(key)
	if err != nil {
		panic("Invalid argument \"" + key + "\"")
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
