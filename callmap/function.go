package callmap

import "github.com/alexd765/jsbundler/ast"

// Function is a function declaration.
type Function struct {
	Calls     []Call
	Functions []Function
	Name      string

	types map[string]struct{}
}

func newFunction(ast *ast.Node) *Function {
	fn := &Function{
		Name: ast.Name,
		types: map[string]struct{}{
			"CallExpression":      struct{}{},
			"FunctionDeclaration": struct{}{},
		},
	}
	fn.walk(ast)
	return fn
}

func (fn *Function) walk(ast *ast.Node) {
	for _, child := range ast.Children {
		nodes := child.WalkTo(fn.types)
		for _, node := range nodes {
			switch node.Type {
			case "CallExpression":
				for _, childNode := range node.Children {
					fn.walk(childNode)
				}
				fn.Calls = append(fn.Calls, Call{Name: node.Name, From:node.From})
			case "FunctionDeclaration":
				fn.Functions = append(fn.Functions, *newFunction(node))
			}
		}
	}
}
