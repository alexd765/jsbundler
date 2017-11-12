package ast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

// ParseFile parses javascript file and returs the AST.
func ParseFile(path string) (*Node, error) {
	out, err := exec.Command("babylon", path).CombinedOutput()
	if err != nil {
		start := bytes.Index(out, []byte("Unexpected token"))
		length := bytes.IndexByte(out[start:], ')') + 1
		return nil, fmt.Errorf("babylon: parsing '%s': '%s'", path, out[start:start+length])
	}

	var n Node
	if err := json.Unmarshal(out, &n); err != nil {
		return nil, err
	}

	return &n, nil
}
