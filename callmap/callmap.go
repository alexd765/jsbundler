package callmap

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// The Callmap stores nothing yet.
type Callmap struct {
	Files map[string]*File `json:"files"`
}

// New returns an ititialized callmap.
func New() *Callmap {
	return &Callmap{
		Files: make(map[string]*File),
	}
}

// Add a javascript file ora directory to the callmap.
func (c *Callmap) Add(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return c.addDir(path)
	}
	f, err := newFile(path)
	if err != nil {
		return err
	}
	c.Files[path] = f
	return nil
}

func (c *Callmap) addDir(path string) error {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		name := fi.Name()
		if strings.HasPrefix(name, ".") && name != "." {
			log.Printf("skipping '%s'", filepath.Join(path, fi.Name()))
			continue
		}
		if ext := filepath.Ext(name); ext != "" && ext != ".js" && ext != ".jsx" {
			log.Printf("skipping '%s'", filepath.Join(path, fi.Name()))
			continue
		}
		if err := c.Add(filepath.Join(path, fi.Name())); err != nil {
			log.Printf("Error parsing %s: %s", filepath.Join(path, fi.Name()), err)
			continue
		}
	}
	return nil
}
