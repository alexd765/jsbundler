package callmap

import "github.com/alexd765/jsbundler/ast"

// Function is a function declaration.
type Function struct {
	Calls     []Call               `json:"calls,omitempty"`
	Functions map[string]*Function `json:"functions,omitempty"`
	Name      string               `json:"name,omitempty"`
}

func newFunction(ast *ast.Node) *Function {
	fn := &Function{
		Name:      ast.Name,
		Functions: make(map[string]*Function),
	}
	fn.walk(ast)
	return fn
}

var typesInFunction = map[string]struct{}{
	"CallExpression":      struct{}{},
	"FunctionDeclaration": struct{}{},
	"ImportDeclaration":   struct{}{},
}

func (fn *Function) walk(ast *ast.Node) {
	for _, child := range ast.Children {
		nodes := child.WalkTo(typesInFunction)
		for _, node := range nodes {
			switch node.Type {
			case "CallExpression":
				for _, childNode := range node.Children {
					fn.walk(childNode)
				}
				fn.Calls = append(fn.Calls, Call{Name: node.Name, From: node.From})
			case "FunctionDeclaration":
				childFn := newFunction(node)
				fn.Functions[childFn.Name] = childFn
			}
		}
	}
}
