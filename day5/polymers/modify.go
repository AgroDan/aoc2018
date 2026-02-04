package polymers

import "strings"

// This will modify the resulting polymer by removing all instances
// of a particular letter. Taking the letter as a rune input, it will
// return a string removing all instances of that letter, capital and
// lowercase

func ModifyPolymer(input string, toRemove rune) string {
	var lower, upper rune
	if isUpperCase(toRemove) {
		upper = toRemove
		lower = toRemove + 32
	} else {
		lower = toRemove
		upper = toRemove - 32
	}

	out := strings.ReplaceAll(input, string(lower), "")
	out = strings.ReplaceAll(out, string(upper), "")
	return out
}
