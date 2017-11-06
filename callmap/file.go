package callmap

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type File struct {
	Program Program
}

type Program struct {
	Body []map[string]interface{} `json:"body"`
}

func newFile(path string) (*File, error) {
	out, err := exec.Command("babylon", path).CombinedOutput()
	if err != nil {
		log.Printf("err: %s", out)
		return nil, err
	}

	var f File
	if err := json.Unmarshal(out, &f); err != nil {
		return nil, err
	}

	for _, node := range f.Program.Body {
		walk(node)
	}

	return &f, nil
}

func walk(node interface{}) {
	n := node.(map[string]interface{})
	switch n["type"] {
	case "AssignmentExpression":
		walk(n["right"])
	case "BinaryExpression":
		walk(n["left"])
		walk(n["right"])
	case "BlockStatement":
		body := n["body"].([]interface{})
		for _, n2 := range body {
			walk(n2)
		}
	case "ForStatement":
		walk(n["init"])
		walk(n["test"])
		walk(n["update"])
		walk(n["body"])
	case "ReturnStatement":
		walk(n["argument"])
	case "ExpressionStatement":
		walk(n["expression"])
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
		arguments := n["arguments"].([]interface{})
		for _, n2 := range arguments {
			walk(n2)
		}
	case "FunctionDeclaration":
		id := n["id"].(map[string]interface{})
		fmt.Printf("%s(){\n", id["name"])
		walk(n["body"])
		fmt.Printf("}\n")
	case "Identifier", "NumericLiteral", "UpdateExpression", "MemberExpression":
	default:
		log.Print(n["type"])
	}
}
