package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexd765/jsbundler/callmap"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Error: Needs javascript source files as parameters.")
	}

	cm := callmap.New()

	for i := 1; i < len(os.Args); i++ {
		if err := cm.Add(os.Args[i]); err != nil {
			log.Fatalf("Error: %s", err)
		}
	}

	fmt.Println("callmap")
	for path, file := range cm.Files {
		fmt.Printf("%s:%+v\n", path, file)
	}
}
