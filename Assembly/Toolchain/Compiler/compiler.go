package compiler

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*

Compiling steps:
1) Load libraries and other config items
2) Parse structure
3) Expand library commands
4) [Variable handling]
5) Generate label addresses (careful: offset, "set" command)
6) Replace labels
7) Output hex file

*/

// PType is the type of parameter that can be passed to a command
type PType int

const (
	// Reg = register
	Reg PType = iota
	// Lit4 = hex literal
	Lit4
	// Ign = Ignored
	Ign
)

// CType is the type of a command
type CType int

const (
	// Cmd = Regular command
	Cmd CType = iota
	// Raw = Insert raw, used after "SET"
	Raw
	// Lbl = Label
	Lbl
)

// program abstractly represents an MCPC program
type program struct {
	commands []command
}

// command is a single elemental instruction of an MCPC program
type command struct {
	commandType CType
	command     string
	args        []string
}

// assemblerProgram is an MCPC program in raw assembler code form
type assemblerProgram struct {
	lines []tokenLine
}

// tokenLine represents a single line of assembler code
type tokenLine struct {
	raw     string
	label   string
	command string
	args    []string
}

// library represents a library that was specified on the command line
type library struct {
	ReplacementTable map[string][]string
}

// Compile transforms a .ma assembly file to a .mb binary
func Compile(file, output string, offset int, libraries []string) {
	fmt.Println("Compiling " + file + " to " + output + "...")
	if offset != 0 {
		fmt.Printf("> Using offset: %X\n", offset)
	}

}

func readFile(path string) []tokenLine {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR: Can't read input file.")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse each line

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
