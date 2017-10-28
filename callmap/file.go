package callmap

import (
	"bytes"
	"io"
	"io/ioutil"
)

// File describes a javascript file.
type File struct {
	functions map[string]*Function
}

func parseFile(path string) (*File, error) {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	file := &File{
		functions: make(map[string]*Function),
	}

	for pos := 0; pos < len(src); pos++ {
		switch src[pos] {
		case '/':
			pos, err = maybeComment(src, pos)
			if err != nil {
				return nil, err
			}
		case 'f':
			pos, err = file.maybeFunction(src, pos)
			if err != nil {
				return nil, err
			}
		}
	}

	return file, nil
}

func maybeComment(src []byte, pos int) (int, error) {
	switch src[pos+1] {
	case '*':
		i := bytes.Index(src[pos+2:], []byte("*/"))
		if i == -1 {
			return 0, io.ErrUnexpectedEOF
		}
		return pos + 2 + i + 1, nil
	case '/':
		i := bytes.IndexByte(src[pos+2:], '\n')
		if i == -1 {
			return 0, io.ErrUnexpectedEOF
		}
		return pos + 2 + i, nil
	}

	return pos, nil
}

func (f *File) maybeFunction(src []byte, pos int) (int, error) {
	if !bytes.HasPrefix(src[pos:], needle) {
		return pos, nil
	}

	fn, err := parseFunction(src, pos)
	if err != nil {
		return 0, err
	}

	f.functions[fn.Name] = fn
	return fn.End, nil
}

func (f File) String() string {
	str := "functions: "
	for _, fn := range f.functions {
		str += fn.String() + " "
	}
	return str
}
