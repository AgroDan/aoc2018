package main

import (
	"day12/plants"
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
	instructions, state, err := plants.Parse(challengeText)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	fmt.Printf("Initial state: %s\n", string(state))

	// make a deep copy of the state so we can use it for part 2
	// because I can't remember if I modify this variable
	stateCopy := make([]rune, len(state))
	copy(stateCopy, state)

	// // state = plants.PadState(state)

	// for i := 0; i < 20; i++ {
	// 	state = plants.StepGeneration(instructions, state)
	// }

	// // Now we need to calculate the sum of the indices of the pots that have plants in them
	// sum := 0
	// for i, pot := range state {
	// 	if pot == '#' {
	// 		sum += i - 2
	// 	}
	// }

	part1Score := plants.ScoreState(instructions, state, 20)
	fmt.Printf("Part 1 score: %d\n", part1Score)

	fmt.Printf("Finding pattern...")
	stableAt, _, diff := plants.FindPattern(instructions, stateCopy)
	fmt.Printf("Pattern stable starting at iteration %d with constant diff %d\n", stableAt, diff)

	// Get the actual sum at the stable point
	sumAtStable := plants.ScoreState(instructions, stateCopy, stableAt)
	
	// Remaining generations after stable point
	remainingIters := 50000000000 - stableAt
	
	// Final answer: sum at stable point + (remaining generations * constant difference)
	finalScore := sumAtStable + (diff * remainingIters)
	fmt.Printf("Part 2 score: %d\n", finalScore)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
