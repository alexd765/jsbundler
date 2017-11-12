package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/alexd765/jsbundler/callmap"
)

func main() {

	v := flag.Bool("v", false, "print the callmap")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatal("Error: Needs javascript source files as parameters.")
	}

	cm := callmap.New()

	for _, arg := range flag.Args() {
		if err := cm.Add(arg); err != nil {
			log.Fatalf("Error: %s", err)
		}
	}

	out, err := json.MarshalIndent(cm, "", "  ")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if *v {
		fmt.Print(string(out))
	}
}
