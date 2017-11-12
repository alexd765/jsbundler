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

// New returns an initialized Callmap.
func New(paths ...string) (*Callmap, error) {
	var filepaths []string

	for _, p := range paths {
		childpaths, err := walkPath(p)
		if err != nil {
			return nil, err
		}
		filepaths = append(filepaths, childpaths...)
	}

	log.Printf("found %d javascript files", len(filepaths))
	files := make(map[string]*File)

	for _, p := range filepaths {
		file, err := newFile(p)
		if err != nil {
			log.Printf("err: %s", err)
			continue
		}
		files[p] = file
	}

	return &Callmap{files}, nil
}

func walkPath(path string) ([]string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return []string{path}, nil
	}

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filepaths []string
	for _, childFi := range fis {
		name := childFi.Name()
		if strings.HasPrefix(name, ".") && name != "." {
			continue
		}
		if ext := filepath.Ext(name); ext != "" && ext != ".js" && ext != ".jsx" {
			continue
		}
		childpaths, err := walkPath(filepath.Join(path, name))
		if err != nil {
			return nil, err
		}
		filepaths = append(filepaths, childpaths...)
	}

	return filepaths, nil
}
