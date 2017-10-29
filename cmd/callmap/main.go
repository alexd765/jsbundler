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
		if err := cm.AddFile(os.Args[i]); err != nil {
			log.Fatalf("Failed to add %s: %s", os.Args[i], err)
		}
	}

	fmt.Println(cm)
}
