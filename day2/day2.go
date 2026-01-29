package main

import (
	"day2/lettercounter"
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

	twoLetters, threeLetters := 0, 0
	for _, line := range lines {
		boxID := lettercounter.NewBoxID(line)
		if boxID.HasExactCount(2) {
			twoLetters++
		}
		if boxID.HasExactCount(3) {
			threeLetters++
		}
	}
	fmt.Printf("Part 1: %d\n", twoLetters*threeLetters)

	for i, line := range lines {
		for j := i + 1; j < len(lines); j++ {
			if lettercounter.CompareBoxIDs(line, lines[j]) == 1 {
				fmt.Printf("Part 2: %s and %s\n", line, lines[j])
				// Now get the common letters
				common := ""
				for k := 0; k < len(line); k++ {
					if line[k] == lines[j][k] {
						common += string(line[k])
					}
				}
				fmt.Printf("Common letters: %s\n", common)
			}
		}
	}

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
