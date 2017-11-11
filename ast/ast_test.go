package ast

import "testing"

func TestParseFile(t *testing.T) {
	f, err := ParseFile("../callmap/testfiles/mult.js")
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	t.Logf("f: %+v", f)
}
