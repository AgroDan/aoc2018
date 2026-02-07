package main

import (
	"day8/licensetree"
	"flag"
	"fmt"
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
	root := licensetree.BuildTree(challengeText)
	fmt.Printf("Root node has %d children and %d metadata entries\n", len(root.Children), len(root.Metadata))
	fmt.Printf("Part 1, Sum of all metadata entries: %d\n", root.SumMetadata())

	// Now get the total according to the challenge
	partTwoTotal := root.Value()
	fmt.Printf("Part 2, total value of the root node: %d\n", partTwoTotal)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
