package main

import (
	"day7/sleighsteps"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	workerPtr := flag.Int("w", 5, "Number of workers for part 2")
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

	inst := sleighsteps.ParseInstructions(lines)
	part1, err := inst.GetFullSteps()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %s\n", part1)

	inst2 := sleighsteps.ParseInstructions(lines)
	part2, tickCount := inst2.GetOrderWithWorkers(*workerPtr)
	fmt.Printf("Part 2: %s, Time taken: %d\n", part2, tickCount)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
