package ast

import (
	"encoding/json"
	"log"
)

// Node is an AST node.
type Node struct {
	Type     string
	Name     string
	Children []*Node
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Node) UnmarshalJSON(b []byte) error {
	var tmp1 struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(b, &tmp1); err != nil {
		return err
	}
	n.Type = tmp1.Type

	switch n.Type {

	case "AssignmentExpression", "BinaryExpression":
		var tmp2 struct {
			Left  *Node `json:"left"`
			Right *Node `json:"right"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Left, tmp2.Right}

	case "BlockStatement", "Program":
		var tmp2 struct {
			Body []*Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = tmp2.Body

	case "CallExpression":
		var tmp2 struct {
			Callee    *Node
			Arguments []*Node `json:"arguments"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = tmp2.Arguments

	case "ExpressionStatement":
		var tmp2 struct {
			Expression *Node `json:"expression"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Expression}

	case "File":
		var tmp2 struct {
			Program *Node `json:"program"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Program}

	case "ForStatement":
		var tmp2 struct {
			Init   *Node `json:"init"`
			Test   *Node `json:"test"`
			Update *Node `json:"update"`
			Body   *Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Init, tmp2.Test, tmp2.Update, tmp2.Body}

	case "FunctionDeclaration":
		var tmp2 struct {
			ID   *Node `json:"id"`
			Body *Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.ID.Name
		n.Children = []*Node{tmp2.Body}

	case "Identifier":
		var tmp2 struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Name

	case "MemberExpression":
		var tmp2 struct {
			Object   *Node `json:"object"`
			Property *Node `json:"property"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Object, tmp2.Property}

	case "ReturnStatement", "UpdateExpression":
		var tmp2 struct {
			Argument *Node `json:"argument"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Argument}

	case "NumericLiteral":

	default:
		log.Printf("unhandled type %s", n.Type)
	}

	return nil
}
