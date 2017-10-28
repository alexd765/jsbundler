package callmap

import (
	"bytes"
	"fmt"
	"io"
)

var (
	needle = []byte("function")
)

// Function describes a javascript function.
type Function struct {
	Name  string
	Start int
	End   int
}

func findFunction(src []byte, offset int) (*Function, error) {
	if offset >= len(src) {
		return nil, io.EOF
	}

	i := bytes.Index(src[offset:], needle)
	if i == -1 {
		return nil, io.EOF
	}

	// Fake Data
	function := Function{
		Start: offset + i,
		Name:  fmt.Sprintf("func %d", offset+i),
		End:   offset + i + 10,
	}
	return &function, nil
}
