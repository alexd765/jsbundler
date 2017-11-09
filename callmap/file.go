package callmap

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type File struct{}

func newFile(path string) (*File, error) {
	log.Printf("adding '%s'", path)
	out, err := exec.Command("babylon", path).CombinedOutput()
	if err != nil {
		log.Printf("err: %s", out)
		return nil, err
	}

	var ast interface{}
	if err := json.Unmarshal(out, &ast); err != nil {
		return nil, err
	}
	walk(ast)

	return nil, nil
}

func walk(node interface{}) {
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
		"NumericLiteral",
		"StringLiteral":

	default:
		log.Print(n["type"])
	}
}
