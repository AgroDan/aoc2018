package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
	"utils"
)

// I'll just do the function here
func directive(l string, state int) int {
	// This will determine the direction we should head
	// and modify the state accordingly
	sign := l[0]
	value := l[1:]
	switch sign {
	case '+':
		convVal, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting value:", err)
			return state
		}
		return state + convVal
	case '-':
		convVal, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting value:", err)
			return state
		}
		return state - convVal
	default:
		fmt.Println("Unknown directive:", l)
		return state
	}
}

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
	state := 0
	for _, line := range lines {
		state = directive(line, state)
	}
	fmt.Printf("Part 1: %d\n", state)

	// Now to find the first repeated frequency
	frequencyMap := make(map[int]bool)
	state = 0
	frequencyMap[state] = true
	found := false
	i := 0
	for !found {
		if i >= len(lines) {
			i = 0
		}
		state = directive(lines[i], state)
		if frequencyMap[state] {
			fmt.Printf("Part 2: First repeated frequency is %d\n", state)
			found = true
		} else {
			frequencyMap[state] = true
		}
		i++
	}

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
