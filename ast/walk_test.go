package ast

import "testing"

func TestWalkTo(t *testing.T) {
	f, err := ParseFile("../callmap/testfiles/mult.js")
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	types := map[string]struct{}{
		"FunctionDeclaration": struct{}{},
		"CallExpression":      struct{}{},
	}

	nodes := f.WalkTo(types)
	if s := len(nodes); s != 3 {
		t.Errorf("want 3 nodes; got %d", s)
	}
	for i, node := range nodes {
		t.Logf("node %d: %+v", i, node)
	}
}
