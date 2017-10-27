package jsbundler

import (
	"fmt"
	"log"

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

// WalkProgram and print things.
func WalkProgram(program *ast.Program) {
	walkDeclarations(program.DeclarationList)
	walkBody(program.Body)
}

func walkDeclarations(declarations []ast.Declaration) {
	for _, decl := range declarations {
		switch d := decl.(type) {
		case *ast.FunctionDeclaration:
			fmt.Printf("function: %s\n", d.Function.Name.Name)
		default:
			log.Printf("unhandled declaration: %T", decl)
		}
	}
}

func walkBody(statements []ast.Statement) {
	for _, statement := range statements {
		switch s := statement.(type) {
		case *ast.ExpressionStatement:
			walkExpression(s.Expression, 0)
		case *ast.EmptyStatement:
		default:
			log.Printf("unhandled statement %T", statement)
		}
	}
}

func walkExpression(ex ast.Expression, depth int) {
	switch e := ex.(type) {
	case *ast.CallExpression:
		walkCallExpression(e, depth)
	case *ast.Identifier:
		fmt.Printf("%d: %s\n", depth, e.Name)
	case *ast.NumberLiteral:
		fmt.Printf("%d: number\n", depth)
	default:
		log.Printf("unhandled expression %T", ex)
	}
}

func walkCallExpression(ce *ast.CallExpression, depth int) {
	fmt.Printf("%d: call \n", depth)
	switch ca := ce.Callee.(type) {
	case *ast.DotExpression:
		walkDotExpression(ca, depth+1)
	case *ast.Identifier:
		fmt.Printf("%d: %s\n", depth+1, ca.Name)
	default:
		fmt.Printf("%d: %T\n", depth+1, ce.Callee)
	}

	fmt.Printf("%d: with\n", depth+1)
	for _, arg := range ce.ArgumentList {
		walkExpression(arg, depth+2)
	}
}

func walkDotExpression(dot *ast.DotExpression, depth int) {
	fmt.Printf("%d: %s.\n", depth, dot.Identifier.Name)
	walkExpression(dot.Left, depth+1)
}
