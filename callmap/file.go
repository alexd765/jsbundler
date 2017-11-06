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

func walk(node map[string]interface{}) {
	switch node["type"] {
	case "AssignmentExpression":
		right := node["right"].(map[string]interface{})
		walk(right)
	case "BinaryExpression":
		left := node["left"].(map[string]interface{})
		walk(left)
		right := node["right"].(map[string]interface{})
		walk(right)
	case "BlockStatement":
		body := node["body"].([]interface{})
		for _, n := range body {
			n2 := n.(map[string]interface{})
			walk(n2)
		}
	case "ForStatement":
		init := node["init"].(map[string]interface{})
		walk(init)
		test := node["test"].(map[string]interface{})
		walk(test)
		update := node["update"].(map[string]interface{})
		walk(update)
		body := node["body"].(map[string]interface{})
		walk(body)
	case "ReturnStatement":
		argument := node["argument"].(map[string]interface{})
		walk(argument)
	case "ExpressionStatement":
		node2 := node["expression"].(map[string]interface{})
		walk(node2)
	case "CallExpression":
		callee := node["callee"].(map[string]interface{})
		if callee["type"] == "MemberExpression" {
			object := callee["object"].(map[string]interface{})
			property := callee["property"].(map[string]interface{})
			fmt.Printf("call %s.%s()\n", object["name"], property["name"])
		}
		if callee["type"] == "Identifier" {
			fmt.Printf("call %s()\n", callee["name"])
		}
		nodes := node["arguments"].([]interface{})
		for _, n := range nodes {
			n2 := n.(map[string]interface{})
			walk(n2)
		}
	case "FunctionDeclaration":
		id := node["id"].(map[string]interface{})
		fmt.Printf("%s(){\n", id["name"])
		body := node["body"].(map[string]interface{})
		walk(body)
		fmt.Printf("}\n")
	case "Identifier", "NumericLiteral", "UpdateExpression", "MemberExpression":
	default:
		log.Print(node["type"])
	}
}
