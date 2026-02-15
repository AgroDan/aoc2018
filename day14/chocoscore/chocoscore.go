package chocoscore

import (
	"fmt"
	"strconv"
)

type Recipes struct {
	Recipes []int // list of "recipe scores"
	Elf1    int   // index of the current recipe for elf 1
	Elf2    int   // index of the current recipe for elf 2
}

// I'm an idiot! The challenge input is NOT the initial recipes
// but rather the number of recipes to generate before we get the
// next 10 recipes. Yikes.
func NewRecipes(start int) *Recipes {
	// just a single line since the input is just a number
	// that's read in as a string, which helps us to separate
	// each digit
	r := Recipes{
		Recipes: []int{},
		Elf1:    0,
		Elf2:    1,
	}
	for _, char := range strconv.Itoa(start) {
		myNum, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		r.Recipes = append(r.Recipes, myNum)
	}
	return &r
}

func (r *Recipes) AddRecipes() {
	// This only iterates once so the idea is to call this as many
	// times as the challenge requires
	elf1Score := r.Recipes[r.Elf1]
	elf2Score := r.Recipes[r.Elf2]
	newRecipe := elf1Score + elf2Score
	if newRecipe >= 10 {
		r.Recipes = append(r.Recipes, 1)
		newRecipe -= 10
	}
	r.Recipes = append(r.Recipes, newRecipe)
	r.Elf1 = (r.Elf1 + elf1Score + 1) % len(r.Recipes)
	r.Elf2 = (r.Elf2 + elf2Score + 1) % len(r.Recipes)
}

func (r *Recipes) GetScores(start int, numRecipes int) string {
	// This will abitrarily get the score based on the starting
	// index and return the number as a string with numRecipes digits
	result := ""

	// This is rife for hitting an OOB error so I'm going to assume
	// the caller is responsible enough to ensure there is enough padding
	for i := 0; i < numRecipes; i++ {
		result += strconv.Itoa(r.Recipes[start+i])
	}
	return result
}

func (r *Recipes) PrintRecipes() {
	for i, score := range r.Recipes {
		switch i {
		case r.Elf1:
			fmt.Printf("(%d) ", score)
		case r.Elf2:
			fmt.Printf("[%d] ", score)
		default:
			fmt.Printf("%d ", score)
		}
	}
	fmt.Println()
}

func (r *Recipes) GetNumRecipes() int {
	return len(r.Recipes)
}

func (r *Recipes) FindPattern(pattern string) int {
	// This is a bit of a brute force approach but it should be fast enough
	// since the pattern is only 6 digits long. We can just check the last
	// 6 digits after each new recipe is added and see if it matches the pattern.
	// easier to loop over a string of ints than one giant integer
	patternDigits := []int{}
	for _, char := range pattern {
		myNum, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		patternDigits = append(patternDigits, myNum)
	}

	// this assumes the recipes have been calculated
	for i := 0; i < len(r.Recipes)-len(patternDigits); i++ {
		match := false
		if r.Recipes[i] == patternDigits[0] {
			match = true
			for j := 1; j < len(patternDigits); j++ {
				if r.Recipes[i+j] != patternDigits[j] {
					match = false
					break
				}
			}
		}
		if match {
			return i
		}
	}
	return -1
}

// I misunderstood what Part 2 was asking for at first. Apparently
// the idea is that the puzzle input is actually the pattern we need
// to find. So since this could be one of those high-computation puzzles
// I think we can probably just check for the pattern once each recipe
// has been added.

func (r *Recipes) FindPatternAfterEachRecipe(pattern string) int {
	patternDigits := []int{}
	for _, char := range pattern {
		myNum, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		patternDigits = append(patternDigits, myNum)
	}

	// ORIGINALLY i just called the AddRecipes() function here, but
	// after consulting ai i realized that because more than one digit
	// is added at a time, I need to account for the possibility that
	// more than one digit can be added if >10, so I need to check
	// twice if that's the case. so i basically just re-did the
	// function here so I could call it more than once
	for {
		elf1Score := r.Recipes[r.Elf1]
		elf2Score := r.Recipes[r.Elf2]
		newRecipe := elf1Score + elf2Score

		// Check after adding first recipe
		if newRecipe >= 10 {
			r.Recipes = append(r.Recipes, 1)
			if patternFound := checkPattern(r, patternDigits); patternFound >= 0 {
				return patternFound
			}
			newRecipe -= 10
		}

		// Check after adding second recipe
		r.Recipes = append(r.Recipes, newRecipe)
		if patternFound := checkPattern(r, patternDigits); patternFound >= 0 {
			return patternFound
		}

		r.Elf1 = (r.Elf1 + elf1Score + 1) % len(r.Recipes)
		r.Elf2 = (r.Elf2 + elf2Score + 1) % len(r.Recipes)
	}
}

// this is a useless function that i added because I didn't
// understand the challenge. BEHOLD MY SHAME
func checkPattern(r *Recipes, patternDigits []int) int {
	if len(r.Recipes) < len(patternDigits) {
		return -1
	}
	// Check if pattern is at the end
	match := true
	for i := 0; i < len(patternDigits); i++ {
		if r.Recipes[len(r.Recipes)-len(patternDigits)+i] != patternDigits[i] {
			match = false
			break
		}
	}
	if match {
		return len(r.Recipes) - len(patternDigits)
	}
	return -1
}

// this is useless because i didn't understand the challenge
// but I'm leaving it here anyway
func (r *Recipes) FindLeftScore(idx, length int) int {
	if idx-length+1 < 0 {
		panic(fmt.Sprintf("Index %d with length %d is out of bounds for recipes of length %d", idx, length, len(r.Recipes)))
	}
	result := 0
	for i := range length {
		result = result*10 + r.Recipes[idx-length+1+i]
	}
	return result
}
