package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/alexd765/jsbundler/callmap"
)

func main() {

	v := flag.Bool("v", false, "print the callmap")
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatal("err: Needs javascript source files as parameters.")
	}

	start := time.Now()

	cm, err := callmap.New(flag.Args()...)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	fmt.Printf("Parsed %d javascript files in %.1f seconds.\n", len(cm.Files), time.Since(start).Seconds())
	var im, fn, ca int
	for _, f := range cm.Files {
		im += len(f.Imports)
		fn += len(f.Functions)
		ca += len(f.Calls)
		for _, childFn := range f.Functions {
			fn2, ca2 := count(childFn)
			fn += fn2
			ca += ca2
		}
	}
	fmt.Printf("Found %d import statements, %d function declarations and %d function calls.\n", im, fn, ca)

	if *v {
		out, err := json.MarshalIndent(cm, "", "  ")
		if err != nil {
			log.Fatalf("err: %s", err)
		}
		fmt.Print(string(out))
	}
}

func count(fn *callmap.Function) (countFn int, countCa int) {
	for _, childFn := range fn.Functions {
		childCountFn, childCountCa := count(childFn)
		countFn += childCountFn
		childCountCa += childCountCa
	}
	countFn += len(fn.Functions)
	countCa += len(fn.Calls)
	return countFn, countCa
}
