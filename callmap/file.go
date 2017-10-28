package callmap

import (
	"io"
	"io/ioutil"
	"sync"
)

// File describes a javascript file.
type File struct {
	mu        *sync.RWMutex
	functions map[string]*Function
}

func parseFile(path string) (*File, error) {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	file := &File{
		mu:        &sync.RWMutex{},
		functions: make(map[string]*Function),
	}

	for {
		function, err := findFunction(src)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		file.mu.Lock()
		file.functions[function.Name] = function
		file.mu.Unlock()
	}

	return file, nil
}
