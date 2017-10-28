package callmap

import (
	"bytes"
	"io"
)

var (
	needle = []byte("function ")
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
	offset += i + len(needle)
	function := Function{Start: offset}

	i = bytes.IndexByte(src[offset:], '(')
	if i == -1 {
		return nil, io.ErrUnexpectedEOF
	}
	function.Name = string(bytes.TrimSpace(src[offset : offset+i]))
	offset += i + 1

	braces := 0
	for ; offset < len(src); offset++ {
		switch src[offset] {
		case '{':
			braces++
		case '}':
			braces--
			if braces == 0 {
				function.End = offset
				return &function, nil
			}
		}
	}
	return nil, io.ErrUnexpectedEOF
}
