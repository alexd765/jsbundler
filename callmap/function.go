package callmap

import (
	"io"
)

// Function describes a javascript function.
type Function struct {
	Name  string
	start int
	end   int
}

func findFunction(src []byte) (*Function, error) {
	return nil, io.EOF
}
