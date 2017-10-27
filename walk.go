package jsbundler

import (
	"fmt"
	"log"
	"strings"

	"github.com/dop251/goja/ast"
)

// WalkProgram and print things.
func WalkProgram(program *ast.Program) {
	walkDeclarations(program.DeclarationList)
	fmt.Println()
	walkBody(program.Body)
	fmt.Printf("\n\n")
}

func walkDeclarations(declarations []ast.Declaration) {
	for _, decl := range declarations {
		switch d := decl.(type) {
		case *ast.FunctionDeclaration:
			fmt.Printf("%s()\n", d.Function.Name.Name)
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
		fmt.Print(e.Name)
	case *ast.NumberLiteral:
		fmt.Print("number")
	default:
		log.Printf("unhandled expression %T", ex)
	}
}

func walkCallExpression(ce *ast.CallExpression, depth int) {
	fmt.Print("call ")
	switch ca := ce.Callee.(type) {
	case *ast.DotExpression:
		walkDotExpression(ca, depth)
	case *ast.Identifier:
		fmt.Print(ca.Name)
	default:
		log.Printf("undhandled calee %T\n", ce.Callee)
	}

	fmt.Print("() with")
	for _, arg := range ce.ArgumentList {
		fmt.Printf("\n%s", strings.Repeat("  ", depth+1))
		walkExpression(arg, depth+1)
	}
}

func walkDotExpression(dot *ast.DotExpression, depth int) {
	fmt.Printf("%s.", dot.Identifier.Name)
	walkExpression(dot.Left, depth+1)
}
