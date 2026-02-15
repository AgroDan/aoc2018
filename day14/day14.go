package main

import (
	"day14/chocoscore"
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	// filePtr := flag.String("f", "input", "Input file if not 'input'")
	numPtr := flag.Int("n", 10, "Number of iterations to cycle through")
	startNumPtr := flag.Int("s", 37, "Starting recipe number seed")
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
	// 	fmt.Println("Fatal:", err)
	// }

	// Insert code here
	chocoRecipes := chocoscore.NewRecipes(*startNumPtr)
	for i := 0; i < *numPtr+10; i++ {
		chocoRecipes.AddRecipes()
	}
	// chocoRecipes.PrintRecipes()
	part1Score := chocoRecipes.GetScores(*numPtr, 10)
	fmt.Printf("Part 1: %s with %d recipes\n", part1Score, chocoRecipes.GetNumRecipes())

	partTwoChocoRecipes := chocoscore.NewRecipes(*startNumPtr)
	patternIdx := partTwoChocoRecipes.FindPatternAfterEachRecipe(strconv.Itoa(*numPtr))
	if patternIdx >= 0 {
		fmt.Printf("Found pattern!\n")
		fmt.Printf("Part 2: %d recipes before pattern\n", patternIdx)
	}

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
