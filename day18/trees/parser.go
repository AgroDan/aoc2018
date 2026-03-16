package trees

import "github.com/AgroDan/aocutils"

// This will read in an input file and generate a runemap
// based on its contents. The runemap can be copied after
// each "tick" or "minute" and return the new runemap
// based on the result of the rules

type Acres struct {
	aocutils.Runemap
}

func NewAcres(lines []string) Acres {
	return Acres{aocutils.NewRunemap(lines)}
}
