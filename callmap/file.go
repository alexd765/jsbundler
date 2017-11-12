package callmap

import "github.com/alexd765/jsbundler/ast"

// File descibes a javascript file.
type File struct {
	Calls     []Call
	Functions []Function
}

func newFile(path string) (*File, error) {
	ast, err := ast.ParseFile(path)
	if err != nil {
		return nil, err
	}

	f := File{}

	types := map[string]struct{}{
		"CallExpression":      struct{}{},
		"FunctionDeclaration": struct{}{},
	}

	nodes := ast.WalkTo(types)
	for _, node := range nodes {
		switch node.Type {
		case "CallExpression":
			f.Calls = append(f.Calls, Call{Name: node.Name})
		case "FunctionDeclaration":
			f.Functions = append(f.Functions, Function{Name: node.Name})
		}
	}

	return &f, nil
}

/*
func walk(node interface{}) {
	if node == nil {
		return
	}

	if nodes, ok := node.([]interface{}); ok {
		for _, n := range nodes {
			walk(n)
		}
		return
	}

	n, ok := node.(map[string]interface{})
	if !ok {
		log.Printf("unexpected node: %+v:", node)
		return
	}
	switch n["type"] {

	case "AssignmentExpression":
		walk(n["left"])
		walk(n["right"])

	case "ArrayExpression":
		walk(n["elements"])

	case "ArrowFunctionExpression":
		walk(n["body"])

	case "AwaitExpression":
		walk(n["argument"])

	case "BinaryExpression":
		walk(n["left"])
		walk(n["right"])

	case "BindExpression":
		walk(n["object"])
		walk(n["callee"])

	case "BlockStatement":
		walk(n["body"])

	case "CallExpression":
		callee := n["callee"].(map[string]interface{})
		if callee["type"] == "MemberExpression" {
			object := callee["object"].(map[string]interface{})
			property := callee["property"].(map[string]interface{})
			fmt.Printf("call %s.%s()\n", object["name"], property["name"])
		}
		if callee["type"] == "Identifier" {
			fmt.Printf("call %s()\n", callee["name"])
		}
		walk(n["arguments"])

	case "CatchClause":
		walk(n["param"])
		walk(n["body"])

	case "Class":
		walk(n["superClass"])
		walk(n["body"])

	case "ConditionalExpression":
		walk(n["test"])
		walk(n["consequent"])
		walk(n["alternate"])

	case "DoExpression":
		walk(n["body"])

	case "DoWhileStatement":
		walk(n["body"])
		walk(n["test"])

	case "ExportAllDeclaration":
		fmt.Print("export *")
		walk(n["declaration"])

	case "ExportDefaultDeclaration":
		fmt.Print("export default ")
		walk(n["declaration"])

	case "ExportNamedDeclaration":
		fmt.Print("export ")
		walk(n["declaration"])

	case "ExpressionStatement":
		walk(n["expression"])

	case "File":
		walk(n["program"])

	case "ForInStatement":
		walk(n["left"])
		walk(n["right"])
		walk(n["body"])

	case "ForStatement":
		walk(n["init"])
		walk(n["test"])
		walk(n["update"])
		walk(n["body"])

	case "FunctionDeclaration":
		id := n["id"].(map[string]interface{})
		fmt.Printf("%s(){\n", id["name"])
		walk(n["body"])
		fmt.Printf("}\n")

	case "ImportDeclaration":
		fmt.Print("import ")
		walk(n["specifiers"])
		source := n["source"].(map[string]interface{})
		fmt.Printf("from %s\n", source["value"])

	case "ImportDefaultSpecifier":
		local := n["local"].(map[string]interface{})
		fmt.Printf("%s ", local["name"])

	case "ImportNamespaceSpecifier":
		local := n["local"].(map[string]interface{})
		fmt.Printf("* as %s ", local["name"])

	case "ImportSpecifier":
		imported := n["imported"].(map[string]interface{})
		fmt.Printf("{%s} ", imported["name"])

	case "IfStatement":
		walk(n["test"])
		walk(n["consequent"])
		walk(n["alternate"])

	case "LabeledStatement":
		walk(n["body"])

	case "LogicalExpression":
		walk(n["left"])
		walk(n["right"])

	case "MemberExpression":
		walk(n["object"])
		walk(n["property"])

	case "NewExpression":
		callee := n["callee"].(map[string]interface{})
		if callee["type"] == "MemberExpression" {
			object := callee["object"].(map[string]interface{})
			property := callee["property"].(map[string]interface{})
			fmt.Printf("call %s.%s()\n", object["name"], property["name"])
		}
		if callee["type"] == "Identifier" {
			fmt.Printf("call %s()\n", callee["name"])
		}
		walk(n["arguments"])

	case "ObjectExpression":
		walk(n["properties"])

	case "ObjectProperty":
		walk(n["value"])

	case "Program":
		walk(n["body"])

	case "ReturnStatement":
		walk(n["argument"])

	case "SequenceExpression":
		walk(n["expressions"])

	case "SpreadElement":
		walk(n["argument"])

	case "SwitchCase":
		walk(n["test"])
		walk(n["consequent"])

	case "SwitchStatement":
		walk(n["discriminant"])
		walk(n["cases"])

	case "TaggedTemplateExpression":
		walk(n["tag"])

	case "TemplateLiteral":
		walk(n["expressions"])

	case "ThrowStatement":
		walk(n["argument"])

	case "TryStatement":
		walk(n["block"])
		walk(n["handler"])
		walk(n["finalizer"])

	case "UnaryExpression":
		walk(n["argument"])

	case "UpdateExpression":
		walk(n["argument"])

	case "VariableDeclaration":
		walk(n["declarations"])

	case "VariableDeclarator":
		walk(n["init"])

	case "WhileStatement":
		walk(n["test"])
		walk(n["body"])

	case "WithStatement":
		walk(n["object"])
		walk(n["body"])

	case "YieldExpression":
		walk(n["argument"])

	case
		"BooleanLiteral",
		"BreakStatement",
		"ContinueStatement",
		"EmptyStatement",
		"ForOfStatement",
		"DebuggerStatement",
		"Identifier",
		"NullLiteral",
		"NumericLiteral",
		"StringLiteral":

	default:
		log.Print(n["type"])
	}
}
*/
