package main

import (
	"day3/fabric"
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

	var fabricSlice []*fabric.Fabric
	for _, l := range lines {
		f := fabric.NewFabric(l)
		fabricSlice = append(fabricSlice, f)
	}

	coordMap := fabric.AllPossibleCoordinates(fabricSlice)
	totalClaimedFabric := 0
	for _, v := range coordMap {
		if v > 1 {
			totalClaimedFabric++
		}
	}

	fmt.Printf("Part 1: %d\n", totalClaimedFabric)

	// Now find the one fabric that doesn't overlap with another claim.

	for _, f := range fabricSlice {
		overlaps := false
		for _, otherF := range fabricSlice {
			if f.ID == otherF.ID {
				continue
			}
			if f.Overlaps(*otherF) {
				overlaps = true
				break
			}
		}
		if !overlaps {
			fmt.Printf("Part 2: Non-overlapping fabric ID is %d\n", f.ID)
			break
		}
	}

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
