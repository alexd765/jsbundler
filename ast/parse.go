package ast

import (
	"encoding/json"
	"log"
	"os/exec"
)

// ParseFile parses javascript file and returs the AST.
func ParseFile(path string) (*Node, error) {
	log.Printf("adding '%s'", path)
	out, err := exec.Command("babylon", path).CombinedOutput()
	if err != nil {
		log.Printf("err: %s", out)
		return nil, err
	}

	var n Node
	if err := json.Unmarshal(out, &n); err != nil {
		return nil, err
	}

	return &n, nil
}
