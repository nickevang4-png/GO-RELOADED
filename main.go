package main

import (
	"log"
	"os"
	"strings"

	"GO-RELOADED/processor"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s <input> <output>", os.Args[0])
	}
	in := os.Args[1]
	out := os.Args[2]

	b, err := os.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}

	// feed processor with lines
	lines := strings.Split(string(b), "\n")
	resLines := processor.Process(lines)

	// join result and write
	res := strings.Join(resLines, "\n")
	if err := os.WriteFile(out, []byte(res), 0o644); err != nil {
		log.Fatal(err)
	}
}
