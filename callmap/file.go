package callmap

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type File struct{}

func newFile(path string) (*File, error) {
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

	n := node.(map[string]interface{})
	switch n["type"] {

	case "AssignmentExpression":
		walk(n["right"])

	case "BinaryExpression":
		walk(n["left"])
		walk(n["right"])

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

	case "ExpressionStatement":
		walk(n["expression"])

	case "File":
		walk(n["program"])

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

	case "Program":
		walk(n["body"])

	case "ReturnStatement":
		walk(n["argument"])

	case "Identifier",
		"MemberExpression",
		"NumericLiteral",
		"UpdateExpression":

	default:
		log.Print(n["type"])
	}
}
