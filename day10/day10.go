package main

import (
	"day10/beams"
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

	beamSet := beams.NewBeamCollection(lines)
	timeElapsed := 0
	for {
		if beamSet.CheckAdjacency() {
			break
		}
		beamSet.Step()
		timeElapsed++
	}
	fmt.Printf("Time elapsed: %d\n", timeElapsed)
	beamSet.Display()

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
