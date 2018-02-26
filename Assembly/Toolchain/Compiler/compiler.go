package compiler

import "fmt"

// Compile transforms a .ma assembly file to a .mb binary
func Compile(file, output string, offset int, libraries []string) {
	fmt.Println("Compiling " + file + " to " + output + "...")
	if offset != 0 {
		fmt.Printf("> Using offset: %X\n", offset)
	}
}
