package main

import (
	"day11/fuelcells"
	"flag"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// filePtr := flag.String("f", "input", "Input file if not 'input'")
	numPtr := flag.Int("n", 1133, "Number of iterations to run if not 1133 (my puzzle input)")
	// any additional flags add here

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	// lines, err := utils.GetFileLines(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here
	cell, power := fuelcells.FindMaxPower(*numPtr)
	fmt.Printf("Part 1: The cell with the most power is %s with a power level of %d\n", cell, power)

	parttwocell, size, power := fuelcells.FindBestPowerLevelPart2(*numPtr)
	fmt.Printf("Part 2: The cell with the most power is %s with a power level of %d and a size of %d\n", parttwocell, power, size)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
