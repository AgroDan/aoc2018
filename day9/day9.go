package main

import (
	"day9/marbles"
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

	partOneGame := marbles.Parse(challengeText)
	partOneGame.Play()
	fmt.Printf("Part One: %d\n", partOneGame.HighScore())

	partTwoGame := marbles.Parse(challengeText)
	partTwoGame.Marbles *= 100
	partTwoGame.Play()
	fmt.Printf("Part Two: %d\n", partTwoGame.HighScore())

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
