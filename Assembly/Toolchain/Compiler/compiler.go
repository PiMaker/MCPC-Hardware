package compiler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
type library []libraryEntry

// libraryEntry is a single replacement instruction loaded from a library
type libraryEntry struct {
	capture     *regexp.Regexp
	replacement string
}

var libraryReplaceeRegex = regexp.MustCompile("- (\\S+) ?(\\S*?)? ?(\\S*?)? ?(\\S*?)? ?=")
var libraryReplacementRegex = regexp.MustCompile("=(.*?)(-\\D|$)")
var paramTypeRegex = regexp.MustCompile("\\.(reg|lit)\\d{0,2}")

// Compile transforms a .ma assembly file to a .mb binary
func Compile(file string, offset int, libraries []string, autoJump bool) []byte {
	fmt.Println("Compiling " + file)
	if offset != 0 {
		fmt.Printf("Using offset: %X (Auto-Jump: %t)\n", offset, autoJump)
	}

	libs := make([]library, len(libraries))
	for i, libPath := range libraries {
		libs[i] = loadLibrary(libPath)
	}

	readFile(file, libs)

	return []byte{0x3, 0x6, 0xff, 0xfe, 0xd3, 0x3b, 0x5, 0x3a, 0x5b, 0x33}
}

func loadLibrary(path string) library {
	fmt.Println("Loading library: " + path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR: Can't read input file.")
		os.Exit(1)
	}
	defer file.Close()

	var lib library

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		// Parse each line
		line := scanner.Text()
		replaceeMatch := libraryReplaceeRegex.FindStringSubmatch(line)
		if len(replaceeMatch) == 0 {
			log.Fatal("Could not load library, parser error on line: " + strconv.Itoa(lineNum))
		}

		// Remove first entry (full match)
		replaceeMatch = replaceeMatch[1:]

		// Remove empty entries
		for i := 0; i < len(replaceeMatch); i++ {
			if replaceeMatch[i] == "" {
				replaceeMatch = removeIndex(replaceeMatch, i)
				i--
			}
		}

		// Generate replacee map
		replaceeMap := make(map[string]string)
		captureString := replaceeMatch[0]

		for i, v := range replaceeMatch[1:] {
			captureString += " (\\S+)"
			replaceeMap[paramTypeRegex.ReplaceAllString(v, "")] = "$" + strconv.Itoa(i+1)
		}

		// Generate replacement
		var replacement []string
		split := strings.Split(strings.Trim(libraryReplacementRegex.FindString(line), " -="), ",")

		for _, v := range split {
			v2 := strings.TrimSpace(v)
			for rk, rv := range replaceeMap {
				v2 = strings.Replace(v2, rk, rv, -1)
			}
			replacement = append(replacement, v2)
		}

		lib = append(lib, libraryEntry{
			capture:     regexp.MustCompile(captureString),
			replacement: strings.Join(replacement, "\n"),
		})

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lib
}

func removeIndex(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

func readFile(path string, libs []library) []tokenLine {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR: Can't read input file.")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse each line
		// DEBUG, TODO: Remove immediate replacing of lib commands
		t := scanner.Text()

		for _, lib := range libs {
			for _, entry := range lib {
				if entry.capture.MatchString(t) {
					t = entry.capture.ReplaceAllString(t, entry.replacement)
				}
			}
		}

		fmt.Println("  " + t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
