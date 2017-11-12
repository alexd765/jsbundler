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
		log.Fatal("err: Needs javascript source files as parameters.")
	}

	cm, err := callmap.New(flag.Args()...)
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	out, err := json.MarshalIndent(cm, "", "  ")
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	if *v {
		fmt.Print(string(out))
	}
}
