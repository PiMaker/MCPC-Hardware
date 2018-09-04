package compiler

import (
	"github.com/davecgh/go-spew/spew"
	"go/parser"
	"go/token"
)

// Compile parses and transforms a .go source file to MCPC assembler code
func Compile(file string) string {

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, file, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	spew.Dump(f)

	return ""
}
