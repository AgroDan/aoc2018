package main

import (
	"day18/trees"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	iterPtr := flag.Int("i", 100, "Number of iterations to run for part 2")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	lines, err := utils.GetFileLines(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here

	acreage := trees.NewAcres(lines)
	for i := 0; i < 10; i++ {
		acreage = acreage.Tick()
	}

	partTwoAcreage := trees.NewAcres(lines)
	fmt.Printf("Resource value for part 1: %d\n", acreage.ResourceValue())

	fmt.Printf("Looking for patterns in part 2...\n")
	offset := trees.CheckForPattern(partTwoAcreage, *iterPtr)
	fmt.Printf("Discovered offset: %d\n", offset)

	// THIS MESSED ME UP! Apparently I've been incrementing from where
	// it left off, not from a new iteration. What a pain
	anotherAcreage := trees.NewAcres(lines)
	for range offset {
		anotherAcreage = anotherAcreage.Tick()
	}
	fmt.Printf("Resource value for part 2: %d\n", anotherAcreage.ResourceValue())

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
