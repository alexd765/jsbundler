package callmap

import (
	"io/ioutil"
	"log"
	"os"
)

// The Callmap stores nothing yet.
type Callmap struct{}

// Add a javascript file ora directory to the callmap.
func (c *Callmap) Add(path string) error {
	log.Printf("adding '%s'", path)
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return c.addDir(path)
	}
	_, err = newFile(path)
	return err
}

func (c *Callmap) addDir(path string) error {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		c.Add(fi.Name())
	}
	return nil
}
