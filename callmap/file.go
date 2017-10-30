package callmap

import (
	"bytes"
	"io"
	"io/ioutil"
)

// File describes a javascript file.
type File struct {
	functions map[string]*Function
	src       []byte
	pos       int
}

func newFile(path string) (*File, error) {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	f := &File{
		functions: make(map[string]*Function),
		src:       src,
	}
	if err := f.parse(); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *File) parse() error {
	for ; f.pos < len(f.src); f.pos++ {
		switch f.src[f.pos] {
		case '/':
			if err := f.maybeComment(); err != nil {
				return err
			}
		case 'f':
			if err := f.maybeFunction(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (f *File) maybeComment() error {
	switch f.src[f.pos+1] {
	case '*':
		i := bytes.Index(f.src[f.pos+2:], []byte("*/"))
		if i == -1 {
			return io.ErrUnexpectedEOF
		}
		f.pos += 2 + i + 1
	case '/':
		i := bytes.IndexByte(f.src[f.pos+2:], '\n')
		if i == -1 {
			return io.ErrUnexpectedEOF
		}
		f.pos += 2 + i
	}

	return nil
}

func (f *File) maybeFunction() error {
	if !bytes.HasPrefix(f.src[f.pos:], needle) {
		return nil
	}

	fn, err := parseFunction(f.src, f.pos)
	if err != nil {
		return err
	}

	f.functions[fn.Name] = fn
	f.pos = fn.End
	return nil
}

func (f File) String() string {
	str := "functions: "
	for _, fn := range f.functions {
		str += fn.String() + " "
	}
	return str
}
