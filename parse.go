package jsbundler

import (
	"github.com/dop251/goja/ast"
	"github.com/dop251/goja/parser"
)

// Parse a javascript string into a program.
func Parse(js string) (*ast.Program, error) {
	program, err := parser.ParseFile(nil, "", js, 0)
	if err != nil {
		return nil, err
	}
	return program, nil
}
