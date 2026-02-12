package plants

import (
	"fmt"
	"strings"
)

// Here will be some helper functions to determine
// pattern stability.

// First, a scoring function
func ScoreState(instructions map[string]rune, state []rune, iter int) int {
	// This will return the score of the state
	// given the amount of iterations we want

	// First, iterate
	for i := 0; i < iter; i++ {
		state = StepGeneration(instructions, state)
	}

	score := 0
	for i, pot := range state {
		if pot == '#' {
			score += i - (2 * iter)
		}
	}
	return score
}

func offsetScore(state []rune, offset int) int {
	// returns the score of a state with an offset,
	// used for pattern recognition. I first went with
	// the idea of the pattern itself being the problem
	// but really should be looking for a score pattern.
	score := 0
	for i, pot := range state {
		if pot == '#' {
			score += i - offset
		}
	}
	return score
}

func depadState(state []rune) string {
	// This will be used for pattern recognition
	// It will remove the padding and return the string of the state
	// It will also return the index of the first plant, so we can calculate the score later
	giganticString := string(state)
	pruned := strings.Trim(giganticString, ".")
	return pruned
}

func FindPattern(instructions map[string]rune, state []rune) (int, int, int) {
	// This will iterate indefinitely until a pattern is found, it will
	// then return the _first_ time it saw that pattern, and when it repeated.
	var seenSums []int
	var diffs []int

	for i := 0; ; i++ {
		state = StepGeneration(instructions, state)
		thisScore := offsetScore(state, 2*(i+1))

		seenSums = append(seenSums, thisScore)

		if i > 0 {
			diffs = append(diffs, thisScore-seenSums[i-1])

			if i > 10 {
				// check the last 5 diffs to see if they are the same
				lastDiffs := diffs[len(diffs)-5:]
				allSame := true
				for j := 1; j < len(lastDiffs); j++ {
					if lastDiffs[j] != lastDiffs[0] {
						allSame = false
						break
					}
				}
				if allSame {
					return i - 4, i, lastDiffs[0]
				}
			}
		} else {
			diffs = append(diffs, 0)
		}
	}
}

func DebugPrintState(instructions map[string]rune, state []rune, iter int) {
	// This just prints out the states for debugging purposes
	var seenPatterns []string

	for i := 0; i < iter; i++ {
		state = StepGeneration(instructions, state)
		pruned := depadState(state)
		seenPatterns = append(seenPatterns, pruned)
	}

	for i, pattern := range seenPatterns {
		fmt.Printf("After generation %02d: %s\n", i+1, pattern)
	}
}
