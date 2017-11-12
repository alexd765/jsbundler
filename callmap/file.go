package callmap

import "github.com/alexd765/jsbundler/ast"

// File descibes a javascript file.
type File struct {
	Calls     []Call
	Functions []Function
	Imports   []Import
}

func newFile(path string) (*File, error) {
	ast, err := ast.ParseFile(path)
	if err != nil {
		return nil, err
	}

	calls, fns, imports := walk(ast)
	f := File{
		Calls:     calls,
		Functions: fns,
		Imports:   imports,
	}

	return &f, nil
}

var types = map[string]struct{}{
	"CallExpression":      struct{}{},
	"FunctionDeclaration": struct{}{},
	"ImportDeclaration":   struct{}{},
}

func walk(ast *ast.Node) ([]Call, []Function, []Import) {
	var calls []Call
	var fns []Function
	var imports []Import

	nodes := ast.WalkTo(types)
	for _, node := range nodes {
		switch node.Type {
		case "CallExpression":
			for _, childNode := range node.Children {
				childCalls, _, _ := walk(childNode)
				calls = append(calls, childCalls...)
			}
			calls = append(calls, Call{Name: node.Name})
		case "FunctionDeclaration":
			fn := Function{Name: node.Name}
			for _, childNode := range node.Children {
				childCalls, childFns, _ := walk(childNode)
				fn.Calls = append(fn.Calls, childCalls...)
				fn.Functions = append(fn.Functions, childFns...)
			}
			fns = append(fns, fn)
		case "ImportDeclaration":
			imports = append(imports, Import{Name: node.Name})
		}
	}

	return calls, fns, imports
}
