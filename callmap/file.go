package callmap

import "github.com/alexd765/jsbundler/ast"

// File descibes a javascript file.
type File struct {
	Calls     []Call
	Functions []Function
	Imports   []Import

	types map[string]struct{}
}

func newFile(path string) (*File, error) {
	ast, err := ast.ParseFile(path)
	if err != nil {
		return nil, err
	}

	f := &File{
		types: map[string]struct{}{
			"CallExpression":      struct{}{},
			"FunctionDeclaration": struct{}{},
			"ImportDeclaration":   struct{}{},
		},
	}
	f.walk(ast)

	return f, nil
}

func (f *File) walk(ast *ast.Node) {
	nodes := ast.WalkTo(f.types)
	for _, node := range nodes {
		switch node.Type {
		case "CallExpression":
			for _, childNode := range node.Children {
				f.walk(childNode)
			}
			f.Calls = append(f.Calls, Call{Name: node.Name, From: node.From})
		case "FunctionDeclaration":
			f.Functions = append(f.Functions, *newFunction(node))
		case "ImportDeclaration":
			f.Imports = append(f.Imports, Import{Name: node.Name, From: node.From})
		}
	}
}
