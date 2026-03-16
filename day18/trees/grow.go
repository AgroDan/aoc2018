package trees

import "github.com/AgroDan/aocutils"

// This will involve all the movements
// and helper functions necessary to
// grow, cut, or set up lumber yards.

func (a *Acres) Decide(point aocutils.Coord) rune {
	// Given a coordinate, this will determine what the
	// new value of this coordinate should be based on what
	// is currently surrounding this coordinate.

	neighbors := point.Neighbors()
	trees := 0
	lumberyards := 0

	for _, neighbor := range neighbors {
		// First, if this neighbor isn't in bounds, skip
		if !a.IsInBounds(neighbor) {
			continue
		}
		// Next, check the value of this neighbor and count
		// the number of trees and lumberyards
		thisItem, err := a.Get(neighbor)
		if err != nil {
			panic(err) // shouldn't happen but whatev
		}
		switch thisItem {
		case '|':
			trees++
		case '#':
			lumberyards++
		}
	}

	// Now apply based on what the current value of this point is
	thisItem, err := a.Get(point)
	if err != nil {
		panic(err) // shouldn't happen but whatev
	}
	switch thisItem {
	case '.':
		if trees >= 3 {
			return '|'
		}
	case '|':
		if lumberyards >= 3 {
			return '#'
		}
	case '#':
		if lumberyards == 0 || trees == 0 {
			return '.'
		}
	}
	return thisItem
}

// I hope go is as good at garbage collection as it claims to be
func (a *Acres) Tick() Acres {
	// This will return a new Acres struct with the new runemap
	// after applying the rules to each coordinate

	paintedCanvas := aocutils.GenerateRunemap(a.Width(), a.Height(), '.')

	for y := 0; y < a.Height(); y++ {
		for x := 0; x < a.Width(); x++ {
			point := aocutils.Coord{X: x, Y: y}
			newVal := a.Decide(point)
			paintedCanvas.Set(point, newVal)
		}
	}

	return Acres{paintedCanvas}
}

// helper function to answer part 1
func (a *Acres) ResourceValue() int {
	// This will count the number of trees and lumberyards and return
	// the product of those two numbers

	trees := 0
	lumberyards := 0

	for y := 0; y < a.Height(); y++ {
		for x := 0; x < a.Width(); x++ {
			point := aocutils.Coord{X: x, Y: y}
			thisItem, err := a.Get(point)
			if err != nil {
				panic(err)
			}
			switch thisItem {
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}
	return trees * lumberyards
}
