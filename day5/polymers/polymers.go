package polymers

import "utils"

// This will mostly just be helper functions from what I can tell.
// Generally what I'll need is to find polarity (based on the case
// of the letter denoting the polymer type), and then remove any
// adjacent polymers that have opposite polarity. So aA will be
// removed, while AA or dJ will not.

// I didn't know I could do this lol:
func isUpperCase(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// Another fun thing I learned, runes are just ASCII values,
// so any lower case letter is always 32 greater than its upper
// case equivalent.
func isSameType(r1, r2 rune) bool {
	if isUpperCase(r1) {
		return r1+32 == r2
	}
	return r1-32 == r2
}

// if two units are the same type, then check if they are opposing
// polarities, and if so then we can remove these two units.
func isOppositePolarity(r1, r2 rune) bool {
	return isSameType(r1, r2) && ((isUpperCase(r1) && !isUpperCase(r2)) || (!isUpperCase(r1) && isUpperCase(r2)))
}

// Need to export _something_, right?
func ReactPolymers(input string) string {
	polymerStack := utils.NewGStack[rune]()

	for _, unit := range input {
		topUnit, ok := polymerStack.Peek()
		if ok && isOppositePolarity(topUnit, unit) {
			// they react, so pop the top unit
			polymerStack.Pop()
		} else {
			// otherwise push the new unit onto the stack
			polymerStack.Push(unit)
		}
	}

	// Now need to build the resulting string from the stack
	resultRunes := make([]rune, 0, polymerStack.Size())
	for !polymerStack.IsEmpty() {
		unit, _ := polymerStack.Pop()
		resultRunes = append(resultRunes, unit)
	}

	// The units will be in reverse order since we popped them off the stack
	// So need to reverse the slice and this is stupid confusing but it works
	// i = 0, j = size of the rune map -1, swap, increment/decrement,
	// meet in the middle
	for i, j := 0, len(resultRunes)-1; i < j; i, j = i+1, j-1 {
		resultRunes[i], resultRunes[j] = resultRunes[j], resultRunes[i]
	}
	return string(resultRunes)
}

// commenting this out because this wasn't the issue, the stack process is very
// efficient already and should hanlde this well enough. I had an off-by-one error
// related to the ingestion of the input data that I've corrected in the utils module
// for my own future sanity.

// // Need to repeat until no more reactions occur, so I'll write a wrapper around the
// // ReactPolymers function that will just keep going until the length stops changing.
// func FullyReactPolymers(input string) string {
// 	previousLength := -1
// 	currentPolymer := input

// 	for previousLength != len(currentPolymer) {
// 		previousLength = len(currentPolymer)
// 		currentPolymer = ReactPolymers(currentPolymer)
// 	}

// 	return currentPolymer
// }
