package main

import (
	"log"
	"os"

	"github.com/alexd765/jsbundler/callmap"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Error: Needs javascript source files as parameters.")
	}

	cm := &callmap.Callmap{}

	for i := 1; i < len(os.Args); i++ {
		cm.AddFile(os.Args[i])
	}
}
