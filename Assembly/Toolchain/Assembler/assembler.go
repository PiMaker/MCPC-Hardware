package assembler

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
	label   []string
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
var charLiteralRegex = regexp.MustCompile("^'.+'$")
var spaceReplaceRegex = regexp.MustCompile("\\'(.*?)\\ (.*?)\\'")
var spaceReplaceDoubleRegex = regexp.MustCompile("\\'\\ \\ \\'")

// Compile transforms a .ma assembly file to a .mb binary
func Compile(file string, offset int, libraries []string, autoJump bool) ([]byte, []byte) {
	log.Println("Compiling " + file)

	sym := make([]byte, 0)

	// Possibly rework this:
	longestDeclaration = 0
	declarationMap = make(map[string]string)

	// Don't allow impossible auto-jump
	if autoJump && offset < 3 {
		autoJump = false
		log.Println("WARNING: Auto-Jump was set, but offset is smaller than 3; Auto-Jump has been disabled")
	}

	if offset > 0 {
		log.Printf("Using offset: %X (Auto-Jump: %t)\n", offset, autoJump)
	}

	if offset < 0 {
		offset = 0
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

						// Handle labels
						replacementTokens[0].label = token.label

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

						////fmt.Println("Replaced \"" + token.raw + "\" with \"" + fmt.Sprintf("%V", replacementTokens) + "\"")

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
		for _, lbl := range token.label {
			labelMap[lbl] = uint16(labelAddr)
			fmt.Println(" > Label " + lbl + " located at 0x" + strconv.FormatInt(int64(labelMap[lbl]), 16))

			// Add to symbol map
			sym = append(sym, []byte(fmt.Sprintf("%04x=%s;", labelMap[lbl], lbl))...)
		}
	}

	// Replace labels
	for _, token := range tokens {
		if token.command == "RAW" && token.raw[0] == '.' {
			addr, ok := labelMap[token.raw]
			if !ok {
				log.Fatalln("ERROR: Undefined label referenced: " + token.raw)
			}
			token.raw = fmt.Sprintf("0x%x", addr)
		}
	}

	// Prepend offset bytes
	if offset > 0 {
		nullCommand := &tokenLine{
			raw:     "0x0",
			command: "RAW",
			label:   []string{},
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
			label:   []string{},
			args:    []string{"SCR1"},
		}
		tokens[1] = &tokenLine{
			raw:     "0x" + strconv.FormatInt(int64(offset), 16),
			command: "RAW",
			label:   []string{},
			args:    make([]string, 0),
		}
		tokens[2] = &tokenLine{
			raw:     "MOV SCR1 PC",
			command: "MOV",
			label:   []string{},
			args:    []string{"SCR1", "PC"},
		}
	}

	// Perform compilation of prepared base symbols to assembly bytes
	output := make([]byte, len(tokens)*2)

	for i, tkn := range tokens {
		////fmt.Println("  COMPILE > " + tkn.raw)
		// Check which base command is used and perform according transform action
		switch tkn.command {
		case "RAW":
			if charLiteralRegex.MatchString(tkn.raw) {
				content := tkn.raw[1 : len(tkn.raw)-1]
				content = strings.Replace(content, "\\n", "\n", -1)
				content = strings.Replace(content, "\\s", " ", -1)
				output[i*2+1] = byte(content[0]) & 0x00FF

				if len(output) > 1 {
					output[i*2] = byte(content[1]) & 0x00FF
				}
			} else {
				n := parseHex(tkn.raw)
				output[i*2] = byte((n & 0xFF00) >> 8)
				output[i*2+1] = byte(n & 0x00FF)
			}

		case "MOV":
			output[i*2] = ParseRegister(tkn.args[1])
			output[i*2+1] = (ParseRegister(tkn.args[0]) << 4) | 0x1
		case "MOVNZ":
			output[i*2] = (ParseRegister(tkn.args[2]) << 4) | ParseRegister(tkn.args[1])
			output[i*2+1] = (ParseRegister(tkn.args[0]) << 4) | 0x2
		case "MOVEZ":
			output[i*2] = (ParseRegister(tkn.args[2]) << 4) | ParseRegister(tkn.args[1])
			output[i*2+1] = (ParseRegister(tkn.args[0]) << 4) | 0x3

		case "BUS":
			output[i*2] = byte(parseHex(tkn.args[1]))
			output[i*2+1] = (ParseRegister(tkn.args[0]) << 4) | 0x4
		case "HOLD":
			output[i*2+1] = 0x5
		case "SET":
			output[i*2] = ParseRegister(tkn.args[0])
			//if output[i*2] == 0xB {
			//	log.Fatalln("ERROR: Cannot SET program counter (PC/0xB)")
			//}
			output[i*2+1] = 0x6

		case "MEMR":
			output[i*2+1] = (ParseRegister(tkn.args[0]) << 4) | 0x5
			output[i*2] = ParseRegister(tkn.args[1])
		case "MEMW":
			output[i*2+1] = (ParseRegister(tkn.args[0]) << 4) | 0x7
			output[i*2] = ParseRegister(tkn.args[1]) << 4

		case "AND", "OR", "XOR", "ADD", "SHFT", "MUL", "GT", "EQ":
			aluCmd(&output, i, tkn)

		case "HALT":
			break

		default:
			log.Println("WARNING: Invalid instruction encountered: \"" + tkn.command + "\" (in \"" + tkn.raw + "\"); Output will be 0x0 (HALT)")
		}
	}

	// Append HALT at end if not already present
	if len(output) > 0 && (output[len(output)-1] != 0 || output[len(output)-2] != 0) {
		output = append(output, []byte{0x0, 0x0}...)
	}

	log.Println("Compilation complete, " + strconv.Itoa(len(output)) + " bytes generated!")

	if len(sym) > 0 {
		return output, sym[:len(sym)-1]
	} else {
		return output, make([]byte, 0)
	}
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
	case "XOR":
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

	out[i*2+1] = ins | (ParseRegister(tkn.args[0]) << 4)
	if tkn.command == "SHFT" {
		if tkn.args[2][0] == '-' {
			// Special care for negative shiftings by manually setting highest bit to 1
			v, _ := strconv.ParseInt(tkn.args[2][3:], 16, 17)
			tkn.args[2] = "0X" + strconv.FormatInt(v|0x8, 16)
		}
		out[i*2] = ParseRegister(tkn.args[1]) | (byte(parseHex(tkn.args[2])&0xF) << 4)
	} else {
		out[i*2] = ParseRegister(tkn.args[1]) | (ParseRegister(tkn.args[2]) << 4)
	}
}

// Parses a hex encoded string with leading "0x" marker to an unsigned 16 bit integer
func parseHex(raw string) uint16 {
	p, _ := strconv.ParseUint(raw[2:], 16, 17)
	return uint16(p)
}

// ParseRegister parses a string representation of a register value to a machine(=MCPC)-readable integer constant
func ParseRegister(reg string) byte {
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
	for strings.HasPrefix(path, "--library=") {
		path = path[len("--library="):] // Weirdness on parameter passing sometimes
	}

	log.Println("Loading library: " + path)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("ERROR: Can't read library file: " + err.Error())
	}
	defer file.Close()

	var lib library

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		// Parse each line
		line := strings.TrimSpace(scanner.Text())
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
		captureString := "(?:\\s|^)" + replaceeMatch[0] // Regex at the beginning takes care that no labels will be replaced

		for i, v := range replaceeMatch[1:] {
			captureString += "\\s+(\\S+)"
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

// Defined globally, not very pretty but gets the job done
var declarationMap map[string]string
var longestDeclaration int

func tokenize(reader io.Reader) []*tokenLine {
	var tokens []*tokenLine

	nextLabel := []string{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// Parse each line
		t := strings.TrimSpace(scanner.Text())

		// Handle comments
		t = strings.TrimSpace(strings.Split(t, ";")[0])
		if t == "" {
			continue
		}

		// Replace spaces in char literals with \s
		t = spaceReplaceDoubleRegex.ReplaceAllString(t, "'\\s\\s'")
		t = spaceReplaceRegex.ReplaceAllString(t, "'$1\\s$2'")

		// Split at spaces
		tspaced := strings.Split(t, " ")
		// Remove empty entries
		for i := 0; i < len(tspaced); i++ {
			if tspaced[i] == "" {
				tspaced = removeIndex(tspaced, i)
				i--
			} else {
				tspaced[i] = strings.TrimSpace(tspaced[i])
				if tspaced[i][0] != '\'' || tspaced[i][len(tspaced[i])-1] != '\'' {
					// Not a char literal, safe to transform to uppercase
					tspaced[i] = strings.ToUpper(tspaced[i])
				}
			}
		}

		// Handle declarations
		if tspaced[0] == "#DECLARE" {
			if len(tspaced) != 3 {
				log.Fatalln("ERROR: Invalid #declare: " + scanner.Text())
			}

			if len(tspaced[2]) < 2 {
				log.Fatalln("ERROR: Invalid #declare (length of replacee has to be at least 2 characters): " + scanner.Text())
			}

			declarationMap[tspaced[2]] = tspaced[1]
			if len(tspaced[2]) > longestDeclaration {
				longestDeclaration = len(tspaced[2])
			}

			continue
		}

		// Label detection
		isLabel := tspaced[0][0] == '.'
		label := []string{}

		if isLabel {
			lineLabel := tspaced[0]

			if len(tspaced) == 1 {
				// Label only, treat as command
				tokens = append(tokens, &tokenLine{
					raw:     lineLabel,
					label:   []string{},
					command: "RAW",
					args:    make([]string, 0),
				})
				nextLabel = []string{}
				continue
			} else if tspaced[1] == "__LABEL_SET" {
				nextLabel = append(nextLabel, lineLabel)
				continue
			}

			tspaced = tspaced[1:]
			label = append(nextLabel, lineLabel)
			nextLabel = []string{}
		} else if nextLabel != nil {
			label = nextLabel
			nextLabel = []string{}
		}

		// Check for raw instructions
		n, err := strconv.ParseInt(t[2:], 16, 17)
		if err == nil {
			// Number literal found
			tokens = append(tokens, &tokenLine{
				raw:     "0x" + strconv.FormatInt(n, 16),
				label:   label,
				command: "RAW",
				args:    make([]string, 0),
			})
			continue
		}
		if t[0] == '\'' && t[len(t)-1] == '\'' {
			// Char literal found
			tokens = append(tokens, &tokenLine{
				raw:     t,
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
				// Special care on iterating the declationMap to allow more complex declarations
				for decLength := longestDeclaration; decLength > 0; decLength-- {
					for k, v := range declarationMap {
						if len(k) == decLength {
							cmdArgs[i] = strings.Replace(cmdArgs[i], k, v, -1)
						}
					}
				}

				if cmdArgs[i][0] == '.' {
					// Append space to allow label handling in library replacing (very hacky haha lmao sorry)
					cmdArgs[i] = cmdArgs[i] + " "
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
