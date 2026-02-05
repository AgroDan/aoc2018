package chronalcoords

import (
	"fmt"
	"utils"
)

// This will parse the listing, ingest them as utils.coords,
// then determine the boundaries based on the min/max x/y values.

type Ccmap struct {
	coords         []utils.Coord
	uX, uY, lX, lY int // upperX, upperY, lowerX, lowerY of boundaries
}

func CreateMapper(lines []string) *Ccmap {
	thisMap := &Ccmap{}
	thisMap.coords = parseCoords(lines)
	thisMap.setBoundaries()
	return thisMap
}

func parseCoords(lines []string) []utils.Coord {
	coords := make([]utils.Coord, len(lines))
	for i, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		coords[i] = utils.Coord{X: x, Y: y}
	}
	return coords
}

// this function just sets the boundaries of all the found coordinates,
// eliminating the need to search for coordinates outside of the "infinite"
// area. Note I'm going to put a border around the edges to ensure I catch
// everything
func (cm *Ccmap) setBoundaries() {
	if len(cm.coords) == 0 {
		return
	}
	cm.uX, cm.lX = cm.coords[0].X, cm.coords[0].X
	cm.uY, cm.lY = cm.coords[0].Y, cm.coords[0].Y

	for _, coord := range cm.coords[1:] {
		if coord.X > cm.uX {
			cm.uX = coord.X
		}
		if coord.X < cm.lX {
			cm.lX = coord.X
		}
		if coord.Y > cm.uY {
			cm.uY = coord.Y
		}
		if coord.Y < cm.lY {
			cm.lY = coord.Y
		}
	}

	// extend the border +1
	cm.uX += 1
	cm.uY += 1
	cm.lX -= 1
	cm.lY -= 1
}

// Just a helper function for me so I can determine how many coordinates
// i can check, if that's even necessary for how I should go about solving
// this.
func (cm *Ccmap) GetPossibleCheckCoords() int {
	totalArea := (cm.uX - cm.lX + 1) * (cm.uY - cm.lY + 1)
	return totalArea - len(cm.coords)
}
