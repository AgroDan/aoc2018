package elfwar

import (
	"math"

	"github.com/AgroDan/aocutils"
)

// Screw it, I'm just going to use BFS for movement. The idea being that
// in all 4 directions, since we're only moving one step at a time, all
// I care about is how many steps to the end point from each of the 4
// points.

func whichStep(start, end aocutils.Coord, b *Battlefield) (aocutils.Coord, int) {
	// This will look at every surrounding [valid] point, and determine
	// which of those points I'll use to move to. This won't be using
	// aStar much to my chagrin, but this will determine the minimum
	// number of steps from each direction, full stop. Then I'll just
	// choose the shortest one, and if there's a tie I'll choose
	// the one in reading order.
	stepsMap := make(map[aocutils.Coord]int)
	directions := start.TrueAllAvailable()
	for _, neighbor := range directions {
		if !isValid(neighbor, b) {
			continue
		}

		// already did this so let's just let'er rip
		steps := bfsHeuristic(neighbor, end, b)
		if steps != -1 {
			stepsMap[neighbor] = steps
		}
	}

	if len(stepsMap) == 0 {
		return start, -1
	}

	// now with a map of all the valid directions, I can find the lowest
	// score. If it's a tie, use reading order
	lowestSteps := math.MaxInt64
	for _, steps := range stepsMap {
		if steps < lowestSteps {
			lowestSteps = steps
		}
	}

	var bestCoords []aocutils.Coord
	for coord, steps := range stepsMap {
		if steps == lowestSteps {
			bestCoords = append(bestCoords, coord)
		}
	}

	readingOrder(&bestCoords)
	return bestCoords[0], lowestSteps
}
