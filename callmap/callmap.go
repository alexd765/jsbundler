package callmap

import (
	"sync"
)

// CallMap is a map which javascript function calls another.
type CallMap struct {
	mu    *sync.RWMutex
	files map[string]*File
}

// New creates an itialized CallMap.
func New() *CallMap {
	return &CallMap{
		mu:    &sync.RWMutex{},
		files: make(map[string]*File),
	}
}

// AddFile adds a javascript file to the CallMap.
func (cm *CallMap) AddFile(path string) error {
	file, err := parseFile(path)
	if err != nil {
		return err
	}

	cm.mu.Lock()
	cm.files[path] = file
	cm.mu.Unlock()

	return nil
}
