package callmap

import (
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

	var offset int
	for {
		function, err := findFunction(src, offset)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		file.functions[function.Name] = function
		offset = function.End
	}

	return file, nil
}

func (f File) String() string {
	str := "functions: "
	for _, fn := range f.functions {
		str += fn.String() + " "
	}
	return str
}
