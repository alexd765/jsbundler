package callmap

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

// The Callmap stores nothing yet.
type Callmap struct {
	Files map[string]*File `json:"files"`
}

// New returns an initialized Callmap.
func New(paths ...string) (*Callmap, error) {
	filepaths := make(chan string, 200)
	eg, _ := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		for _, p := range paths {
			if err := walkPath(p, filepaths); err != nil {
				return err
			}
		}
		close(filepaths)
		return nil
	})

	mu := &sync.Mutex{}
	files := make(map[string]*File)

	for i := 0; i < 8; i++ {
		eg.Go(func() error {
			for p := range filepaths {
				file, err := newFile(p)
				if err != nil {
					log.Printf("err: %s", err)
					continue
				}
				mu.Lock()
				files[p] = file
				mu.Unlock()
			}
			return nil
		})
	}
	eg.Wait()
	return &Callmap{files}, nil
}

func walkPath(path string, filepaths chan string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		filepaths <- path
		return nil
	}

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, childFi := range fis {
		name := childFi.Name()
		if strings.HasPrefix(name, ".") && name != "." {
			continue
		}
		if ext := filepath.Ext(name); ext != "" && ext != ".js" && ext != ".jsx" {
			continue
		}
		if err := walkPath(filepath.Join(path, name), filepaths); err != nil {
			return err
		}
	}

	return nil
}
