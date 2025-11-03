package main

import (
	"fmt"
	"log"
	"os"

	"GO-RELOADED/internal/textops"
	"GO-RELOADED/internal/utils"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go-reloaded <input_file> <output_file>")
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	content, err := utils.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	result := textops.Process(content)

	if err := utils.WriteFile(outputPath, result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ File processed and saved to:", outputPath)
}
