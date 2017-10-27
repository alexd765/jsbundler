package jsbundler

import (
	"testing"
)

func TestParse(t *testing.T) {
	js := `
		function plus(a,b) {
			return a+b;
		}
		function mul(a,b) {
			return a*b;
		}
		log.console(plus(1,2))
	 `

	program, err := Parse(js)
	if err != nil {
		t.Fatal(err)
	}
	WalkProgram(program)
}
