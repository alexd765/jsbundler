package ast

import (
	"encoding/json"
	"log"
)

// Node is an AST node.
type Node struct {
	Type     string
	Name     string
	From     string
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

	case "AssignmentExpression", "AssignmentPattern", "BinaryExpression", "LogicalExpression":
		var tmp2 struct {
			Left  *Node `json:"left"`
			Right *Node `json:"right"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Left, tmp2.Right}

	case "ArrayExpression", "ArrayPattern":
		var tmp2 struct {
			Elements []*Node `json:"elements"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = tmp2.Elements

	case "ArrowFunctionExpression":
		var tmp2 struct {
			Params []*Node `json:"params"`
			Body   *Node   `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = append(tmp2.Params, tmp2.Body)

	case "BlockStatement", "ClassBody", "DoExpression", "LabeledStatement", "Program":
		var tmp2 struct {
			Body []*Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = tmp2.Body

	case "BindExpression":
		var tmp2 struct {
			Callee *Node `json:"callee"`
			Object *Node `json:"object"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Callee, tmp2.Object}

	case "CallExpression", "NewExpression":
		var tmp2 struct {
			Callee    *Node
			Arguments []*Node `json:"arguments"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Callee.Name
		n.From = tmp2.Callee.From
		n.Children = tmp2.Arguments

	case "ClassDeclaration", "ClassExpression":
		var tmp2 struct {
			SuperClass *Node `json:"superClass"`
			Body       *Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.SuperClass, tmp2.Body}

	case "ClassMethod":
		var tmp2 struct {
			Key  *Node `json:"key"`
			Body *Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Key.Name
		n.Children = []*Node{tmp2.Body}

	case "ClassProperty":
		var tmp2 struct {
			Key   *Node `json:"key"`
			Value *Node `json:"value"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Key.Name
		n.Children = []*Node{tmp2.Value}

	case "CatchClause":
		var tmp2 struct {
			Param *Node `json:"param"`
			Body  *Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Param, tmp2.Body}

	case "DoWhileStatement", "WhileStatement":
		var tmp2 struct {
			Body *Node `json:"body"`
			Test *Node `json:"test"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Body, tmp2.Test}

	case "ExportAllDeclaration", "ExportDefaultDeclaration", "ExportNamedDeclaration":
		var tmp2 struct {
			Declaration *Node `json:"declaration"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Declaration}

	case "ExpressionStatement", "JSXExpressionContainer":
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

	case "ForInStatement":
		var tmp2 struct {
			Left  *Node `json:"left"`
			Right *Node `json:"right"`
			Body  *Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Left, tmp2.Right, tmp2.Body}

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

	case "FunctionDeclaration", "ObjectMethod":
		var tmp2 struct {
			ID   *Node `json:"id"`
			Body *Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		if tmp2.ID != nil {
			n.Name = tmp2.ID.Name
		}
		n.Children = []*Node{tmp2.Body}

	case "FunctionExpression":
		var tmp2 struct {
			Params []*Node `json:"params"`
			Body   *Node   `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = append(tmp2.Params, tmp2.Body)

	case "Identifier":
		var tmp2 struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Name

	case "ConditionalExpression", "IfStatement":
		var tmp2 struct {
			Test       *Node `json:"test"`
			Consequent *Node `json:"consequent"`
			Alternate  *Node `json:"alternate"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Test, tmp2.Consequent}
		if tmp2.Alternate != nil {
			n.Children = append(n.Children, tmp2.Alternate)
		}

	case "ImportDeclaration":
		var tmp2 struct {
			Specifiers []*Node `json:"specifiers"`
			Source     *Node   `json:"source"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		for _, spec := range tmp2.Specifiers {
			n.Name += spec.Name + " "
		}
		n.From = tmp2.Source.Name

	case "ImportSpecifier":
		var tmp2 struct {
			Imported *Node
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Imported.Name

	case "ImportDefaultSpecifier", "ImportNamespaceSpecifier":
		var tmp2 struct {
			Local *Node `json:"local"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Local.Name

	case "JSXElement":
		var tmp2 struct {
			Children []*Node `json:"children"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = tmp2.Children

	case "MemberExpression":
		var tmp2 struct {
			Object   *Node `json:"object"`
			Property *Node `json:"property"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Property.Name
		n.From = tmp2.Object.Name
		n.Children = []*Node{tmp2.Object, tmp2.Property}

	case "AwaitExpression", "ReturnStatement", "RestElement", "SpreadElement", "SpreadProperty", "ThrowStatement", "UnaryExpression", "UpdateExpression", "YieldExpression":
		var tmp2 struct {
			Argument *Node `json:"argument"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Argument}

	case "ObjectExpression", "ObjectPattern":
		var tmp2 struct {
			Properties []*Node `json:"properties"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = tmp2.Properties

	case "ObjectProperty":
		var tmp2 struct {
			Value interface{} `json:"value"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		if tmp2.Value != nil {
			if v, ok := tmp2.Value.(string); ok {
				n.Name = v
			}
		}

	case "StringLiteral":
		var tmp2 struct {
			Value string `json:"value"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Name = tmp2.Value

	case "SequenceExpression", "TemplateLiteral":
		var tmp2 struct {
			Expressions []*Node `json:"expressions"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = tmp2.Expressions

	case "SwitchCase":
		var tmp2 struct {
			Test       *Node   `json:"test"`
			Consequent []*Node `json:"conseqeuent"`
		}

		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		if tmp2.Test != nil {
			n.Children = []*Node{tmp2.Test}
		}
		n.Children = append(n.Children, tmp2.Consequent...)

	case "SwitchStatement":
		var tmp2 struct {
			Discriminant *Node   `json:"discriminant"`
			Cases        []*Node `json:"cases"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Discriminant}
		n.Children = append(n.Children, tmp2.Cases...)

	case "TaggedTemplateExpression":
		var tmp2 struct {
			Tag *Node `json:"tag"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Tag}

	case "TryStatement":
		var tmp2 struct {
			Block     *Node `json:"block"`
			Handler   *Node `json:"handler"`
			Finalizer *Node `json:"finalizer"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Block, tmp2.Handler, tmp2.Finalizer}

	case "VariableDeclaration":
		var tmp2 struct {
			Declarations []*Node `json:"declarations"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = tmp2.Declarations

	case "VariableDeclarator":
		var tmp2 struct {
			Init *Node `json:"init"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Init}

	case "WithStatement":
		var tmp2 struct {
			Object *Node `json:"object"`
			Body   *Node `json:"body"`
		}
		if err := json.Unmarshal(b, &tmp2); err != nil {
			return err
		}
		n.Children = []*Node{tmp2.Object, tmp2.Body}

	case
		"BooleanLiteral",
		"BreakStatement",
		"ContinueStatement",
		"EmptyStatement",
		"ForOfStatement",
		"DebuggerStatement",
		"NullLiteral",
		"NumericLiteral",
		"TypeAlias",
		"ThisExpression",
		"JSXText",
		"JSXEmptyExpression",
		"DeclareVariable",
		"RegExpLiteral",
		"InterfaceDeclaration",
		"TypeCastExpression",
		"Super":

	default:
		log.Printf("unhandled type %s", n.Type)
	}

	return nil
}
