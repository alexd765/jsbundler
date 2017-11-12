package callmap

import "github.com/alexd765/jsbundler/ast"

// File descibes a javascript file.
type File struct {
	Calls     []Call
	Functions map[string]*Function
	Imports   []Import
}

func newFile(path string) (*File, error) {
	ast, err := ast.ParseFile(path)
	if err != nil {
		return nil, err
	}

	f := &File{
		Functions: make(map[string]*Function),
	}
	f.walk(ast)

	return f, nil
}

var typesInFile = map[string]struct{}{
	"CallExpression":      struct{}{},
	"FunctionDeclaration": struct{}{},
	"ImportDeclaration":   struct{}{},
}

func (f *File) walk(ast *ast.Node) {
	nodes := ast.WalkTo(typesInFile)
	for _, node := range nodes {
		switch node.Type {
		case "CallExpression":
			for _, childNode := range node.Children {
				f.walk(childNode)
			}
			f.Calls = append(f.Calls, Call{Name: node.Name, From: node.From})
		case "FunctionDeclaration":
			fn := newFunction(node)
			f.Functions[fn.Name] = fn
		case "ImportDeclaration":
			f.Imports = append(f.Imports, Import{Name: node.Name, From: node.From})
		}
	}
}
