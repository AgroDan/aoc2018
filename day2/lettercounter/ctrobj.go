package lettercounter

import "utils"

type BoxID struct {
	Pattern map[rune]int
}

func NewBoxID(id string) *BoxID {
	pattern := make(map[rune]int)
	for _, char := range id {
		pattern[char]++
	}
	return &BoxID{Pattern: pattern}
}

func (b *BoxID) HasExactCount(count int) bool {
	for _, v := range b.Pattern {
		if v == count {
			return true
		}
	}
	return false
}

func CompareBoxIDs(a, b string) int {
	diffCount := 0
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	for i := 0; i < minLength; i++ {
		if a[i] != b[i] {
			diffCount++
		}
	}
	// Shouldn't matter because i'm pretty sure they're the same length but whatever
	return diffCount + utils.Abs(len(a)-len(b))

}
