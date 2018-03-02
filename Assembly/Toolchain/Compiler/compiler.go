package compiler

import (
	"bufio"
	"fmt"
	"io"
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
7) Actually compile prepared commands to assembly
8) Output assembly bytes

*/

// command is a single elemental instruction of an MCPC program
type command struct {
	command string
	args    []string
	isRaw   bool
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
	log.Println("Compiling " + file)

	// Don't allow impossible auto-jump
	if autoJump && offset < 3 {
		autoJump = false
		log.Println("WARNING: Auto-Jump was set, but offset is smaller than 3; Auto-Jump has been disabled")
	}

	if offset != 0 {
		log.Printf("Using offset: %X (Auto-Jump: %t)\n", offset, autoJump)
	}

	// Load libraries
	libs := make([]library, len(libraries))
	for i, libPath := range libraries {
		libs[i] = loadLibrary(libPath)
	}

	// Read and parse source file
	log.Println("Tokenizing...")
	tokens := readFile(file)

	log.Println("Applying library transforms...")
	// Handle each library in a loop until no more replacements have occured
	replaced := 1
	for replaced > 0 {
		replaced = 0

		// Actually process libraries
		for _, lib := range libs {
			for i := 0; i < len(tokens); i++ {
				token := tokens[i]
				for _, r := range lib {
					if r.capture.MatchString(token.raw) {
						rawLibReplacement := r.capture.ReplaceAllString(token.raw, r.replacement)
						replacementTokens := tokenize(strings.NewReader(rawLibReplacement))

						// Perform insert

						// Grow the slice
						tokens = append(tokens, make([]*tokenLine, len(replacementTokens)-1)...)
						// Use copy to move the upper part of the slice out of the way and open a hole
						copy(tokens[i+len(replacementTokens)-1:], tokens[i:])
						// Store the new values
						for ir := 0; ir < len(replacementTokens); ir++ {
							tokens[i+ir] = replacementTokens[ir]
						}

						// Update index
						i += len(replacementTokens) - 1

						replaced++
					}
				}
			}
		}
	}

	log.Println("Parsing labels...")

	// Parse labels
	labelMap := make(map[string]uint16)
	for labelAddr, token := range tokens {
		if token.label != "" {
			labelMap[token.label] = uint16(labelAddr) - uint16(1)
		}
	}

	// Replace labels
	for _, token := range tokens {
		if token.command == "RAW" && token.raw[0] == '.' {
			token.raw = fmt.Sprintf("0x%x", labelMap[token.raw])
		}
	}

	// Prepend offset bytes
	if offset > 0 {
		nullCommand := &tokenLine{
			raw:     "0x0",
			command: "RAW",
			label:   "",
			args:    make([]string, 0),
		}
		offsetLines := make([]*tokenLine, offset)
		for i := range offsetLines {
			offsetLines[i] = nullCommand
		}
		tokens = append(offsetLines, tokens...)
	}

	// Auto-Jump
	if autoJump {
		tokens[0] = &tokenLine{
			raw:     "SET SCR1",
			command: "SET",
			label:   "",
			args:    []string{"SCR1"},
		}
		tokens[1] = &tokenLine{
			raw:     "0x" + strconv.FormatInt(int64(offset), 16),
			command: "RAW",
			label:   "",
			args:    make([]string, 0),
		}
		tokens[2] = &tokenLine{
			raw:     "MOV SCR1 PC",
			command: "MOV",
			label:   "",
			args:    []string{"SCR1", "PC"},
		}
	}

	// Perform compilation of prepared base symbols to assembly bytes
	output := make([]byte, len(tokens)*2)

	for i, tkn := range tokens {
		// Check which base command is used and perform according transform action
		switch tkn.command {
		case "RAW":
			n := parseHex(tkn.raw)
			output[i*2] = byte((n & 0xFF00) >> 8)
			output[i*2+1] = byte(n & 0x00FF)
		case "MOV":
			output[i*2] = parseRegister(tkn.args[1])
			output[i*2+1] = (parseRegister(tkn.args[0]) << 4) | 0x1
		case "MOVNZ":
			output[i*2] = (parseRegister(tkn.args[2]) << 4) | parseRegister(tkn.args[1])
			output[i*2+1] = (parseRegister(tkn.args[0]) << 4) | 0x2
		case "MOVEZ":
			output[i*2] = (parseRegister(tkn.args[2]) << 4) | parseRegister(tkn.args[1])
			output[i*2+1] = (parseRegister(tkn.args[0]) << 4) | 0x3
		case "BUS":
			output[i*2] = byte(parseHex(tkn.args[1]))
			output[i*2+1] = (parseRegister(tkn.args[0]) << 4) | 0x4
		case "HOLD":
			output[i*2+1] = 0x5
		case "SET":
			output[i*2] = parseRegister(tkn.args[0])
			output[i*2+1] = 0x6
		case "AND", "OR", "NOT", "ADD", "SHFT", "MUL", "GT", "EQ":
			aluCmd(&output, i, tkn)
		case "HALT":
		default:
			log.Println("WARNING: Invalid instruction encountered: \"" + tkn.command + "\" (in \"" + tkn.raw + "\"); Output will be 0x0 (HALT)")
		}
	}

	log.Println("Compilation complete, " + strconv.Itoa(len(output)) + " bytes generated!")

	return output
}

// Transforms an ALU command token to assembly
func aluCmd(output *[]byte, i int, tkn *tokenLine) {
	out := *output

	var ins byte
	switch tkn.command {
	case "AND":
		ins = 0x8
	case "OR":
		ins = 0x9
	case "NOT":
		ins = 0xA
	case "ADD":
		ins = 0xB
	case "SHFT":
		ins = 0xC
	case "MUL":
		ins = 0xD
	case "GT":
		ins = 0xE
	case "EQ":
		ins = 0xF
	}

	out[i*2+1] = ins | (parseRegister(tkn.args[0]) << 4)
	if tkn.command == "SHFT" {
		if tkn.args[2][0] == '-' {
			// Special care for negative shiftings by manually setting highest bit to 1
			v, _ := strconv.ParseInt(tkn.args[2][3:], 16, 16)
			tkn.args[2] = strconv.FormatInt(v|0x8, 16)
		}
		out[i*2] = parseRegister(tkn.args[1]) | (byte(parseHex(tkn.args[2])&0xF) << 4)
	} else {
		out[i*2] = parseRegister(tkn.args[1]) | (parseRegister(tkn.args[2]) << 4)
	}
}

// Parses a hex encoded string with leading "0x" marker to an unsigned 16 bit integer
func parseHex(raw string) uint16 {
	p, _ := strconv.ParseUint(raw[2:], 16, 16)
	return uint16(p)
}

// Parses a string representation of a register value to a machine(=MCPC)-readable integer constant
func parseRegister(reg string) byte {
	switch reg {
	case "A":
		return 0x0
	case "B":
		return 0x1
	case "C":
		return 0x2
	case "D":
		return 0x3
	case "E":
		return 0x4
	case "F":
		return 0x5
	case "G":
		return 0x6
	case "H":
		return 0x7
	case "SCR1":
		return 0x8
	case "SCR2":
		return 0x9
	case "SP":
		return 0xA
	case "PC":
		return 0xB
	case "0":
		return 0xC
	case "1":
		return 0xD
	case "-1":
		return 0xE
	case "BUS":
		return 0xF
	default:
		log.Println("WARNING: Invalid register name encountered: " + reg + "; Output will be 0x8 (SCR1)")
		// SCR1 is default
		return 0x8
	}
}

func loadLibrary(path string) library {
	log.Println("Loading library: " + path)

	file, err := os.Open(path)
	if err != nil {
		log.Println("ERROR: Can't read input file.")
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
			log.Fatalln("ERROR: Could not load library, parser error on line: " + strconv.Itoa(lineNum))
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
			replaceeMap[":"+paramTypeRegex.ReplaceAllString(v, "")] = "$" + strconv.Itoa(i+1)
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

		//fmt.Println("Lib entry loaded: " + captureString + " transforms to " + strings.Join(replacement, ", "))

		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return lib
}

func removeIndex(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

func readFile(path string) []*tokenLine {
	file, err := os.Open(path)
	if err != nil {
		log.Println("ERROR: Can't read input file.")
		os.Exit(1)
	}
	defer file.Close()

	return tokenize(file)
}

func tokenize(reader io.Reader) []*tokenLine {
	var tokens []*tokenLine

	declarationMap := make(map[string]string)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// Parse each line
		t := strings.TrimSpace(scanner.Text())

		// Handle comments
		t = strings.TrimSpace(strings.Split(t, ";")[0])
		if t == "" {
			continue
		}

		// Split at spaces
		tspaced := strings.Split(t, " ")
		// Remove empty entries
		for i := 0; i < len(tspaced); i++ {
			if tspaced[i] == "" {
				tspaced = removeIndex(tspaced, i)
				i--
			} else {
				tspaced[i] = strings.TrimSpace(tspaced[i])
			}
		}

		// Handle declarations
		if tspaced[0] == "#declare" {
			if len(tspaced) != 3 {
				log.Fatalln("ERROR: Invalid #declare: " + scanner.Text())
			}

			declarationMap[tspaced[2]] = tspaced[1]

			continue
		}

		// Label detection
		isLabel := tspaced[0][0] == '.'
		label := ""
		if isLabel {
			label = tspaced[0]

			if len(tspaced) == 1 {
				// Label only, treat as command
				tokens = append(tokens, &tokenLine{
					raw:     label,
					label:   "",
					command: "RAW",
					args:    make([]string, 0),
				})
				continue
			}

			tspaced = tspaced[1:]
		}

		// Check for raw instructions
		n, err := strconv.ParseInt(t[2:], 16, 16)
		if err == nil {
			// Literal found
			tokens = append(tokens, &tokenLine{
				raw:     "0x" + strconv.FormatInt(n, 16),
				label:   label,
				command: "RAW",
				args:    make([]string, 0),
			})
			continue
		}

		// Process command args
		var cmdArgs []string
		if len(tspaced) > 1 {
			cmdArgs = tspaced[1:]
			for i := 0; i < len(cmdArgs); i++ {
				for k, v := range declarationMap {
					cmdArgs[i] = strings.Replace(cmdArgs[i], k, v, -1)
				}
			}
		}

		// Create and add token
		tokens = append(tokens, &tokenLine{
			raw:     strings.Join(tspaced, " "),
			label:   label,
			command: tspaced[0],
			args:    cmdArgs,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return tokens
}
