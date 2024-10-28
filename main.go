package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("globbing: %v", os.Args[1])

	results, err := DoubleStarGlob(os.Args[1])
	if err != nil {
		log.Fatalf("error: %w", err)
	}

	log.Printf("found %v results", len(results))
}
