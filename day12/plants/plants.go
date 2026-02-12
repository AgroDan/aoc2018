package plants

import (
	"fmt"
	"strings"
)

// This time I'm going to ingest every "plant state" instruction
// as a string and put them as the key value to a map. The value
// of the map will be the state of the 2nd (indexing from 0)
// character in the instruction string. The instruction string will
// be a slice of runes.

// So let's parse as one giant blob of text

// This looks a LOT like AOC 2022 day 17. Part two is basically
// one of those "now try it a gigantic number of times" problems,
// so the way I foresee this working is looking for pattern stabilization
// here. If I can find a pattern that repeats, I can calculate the sum
// number of that pattern and then add the remaining generations to get
// the final sum. What a ride this is gonna be.

func Parse(input string) (map[string]rune, []rune, error) {
	// first split the blob at the double \n
	parts := strings.Split(input, "\n\n")

	// first part is the initial state,
	// but lol make sure you start counting after "Initial state: "
	initialState := []rune(parts[0][15:])

	// second part is the plant state instructions
	instructions := make(map[string]rune)

	// Remember we might have a freakin trailing newline

	lines := strings.TrimSpace(parts[1])
	for _, line := range strings.Split(lines, "\n") {
		// split the line at " => "
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid instruction: %s", line)
		}
		instructions[parts[0]] = []rune(parts[1])[0]
	}

	return instructions, initialState, nil
}

func PadState(state []rune) []rune {
	paddedState := make([]rune, len(state)+4)
	copy(paddedState[2:], state)

	// Also the padding should be .'s, so add these
	for i := 0; i < 2; i++ {
		paddedState[i] = '.'
		paddedState[len(paddedState)-1-i] = '.'
	}

	return paddedState
}

// Now let's simulate a generation
func StepGeneration(instructions map[string]rune, state []rune) []rune {
	// now I'll loop through every 5 character slice of the padded state
	// and check to see if there are any instructions for that current
	// slice. If so, I'll update the state of the 2nd character in the
	// slice. This can grow substantially if part 2 is a gigantic time span
	state = PadState(state)
	newState := make([]rune, len(state))
	copy(newState, state)

	for i := 0; i < len(state)-4; i++ {
		pots := string(state[i : i+5])
		if inst, ok := instructions[pots]; ok {
			newState[i+2] = inst
		} else {
			newState[i+2] = '.'
		}
	}

	return newState
}
