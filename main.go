package main

import (
	"log"
	"os"

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

	res := processor.Process(string(b))

	if err := os.WriteFile(out, []byte(res), 0o644); err != nil {
		log.Fatal(err)
	}
}
a