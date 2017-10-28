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

func parseFunction(src []byte, pos int) (*Function, error) {
	pos += len(needle)
	fn := Function{Start: pos}

	i := bytes.IndexByte(src[pos:], '(')
	if i == -1 {
		return nil, io.ErrUnexpectedEOF
	}
	fn.Name = string(bytes.TrimSpace(src[pos : pos+i]))
	pos += i + 1

	braces := 0
	for ; pos < len(src); pos++ {
		switch src[pos] {
		case '{':
			braces++
		case '}':
			braces--
			if braces == 0 {
				fn.End = pos
				return &fn, nil
			}
		}
	}
	return nil, io.ErrUnexpectedEOF
}

func (fn Function) String() string {
	return fn.Name + "()"
}
