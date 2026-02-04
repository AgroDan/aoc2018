package main

import (
	"day5/polymers"
	"flag"
	"fmt"
	"strings"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	// lines, err := utils.GetFileLines(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// giant text blob:
	challengeText, err := utils.GetTextBlob(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// Insert code here
	resultingPolymer := polymers.ReactPolymers(challengeText)
	resultingPolymer = strings.TrimSpace(resultingPolymer)

	fmt.Printf("Resulting polymer: %s\n", resultingPolymer)
	fmt.Printf("Part 1 answer: %d\n", len(resultingPolymer))

	polymerCheck := "abcdefghijklmnopqrstuvwxyz"
	shortestLength := len(resultingPolymer)

	for _, r := range polymerCheck {
		modifiedPolymer := polymers.ModifyPolymer(resultingPolymer, r)
		reactedPolymer := polymers.ReactPolymers(modifiedPolymer)
		// probably unnecessary but whatever
		reactedPolymer = strings.TrimSpace(reactedPolymer)
		if len(reactedPolymer) < shortestLength {
			shortestLength = len(reactedPolymer)
		}
	}

	fmt.Printf("Part 2 answer: %d\n", shortestLength)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
