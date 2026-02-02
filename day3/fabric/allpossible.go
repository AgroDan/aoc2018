package fabric

import "utils"

// This will be a bit greedy, but I'm going to define here
// a finite set of all possible coordinates from each fabric claim.
// Then I'm going to just count how many of these coordinates show up
// here.

func AllPossibleCoordinates(f []*Fabric) map[utils.Coord]int {
	coordMap := make(map[utils.Coord]int)
	for _, fabric := range f {
		for x := fabric.Pos.X; x < fabric.Pos.X+fabric.Width; x++ {
			for y := fabric.Pos.Y; y < fabric.Pos.Y+fabric.Height; y++ {
				coord := utils.Coord{X: x, Y: y}
				coordMap[coord]++
			}
		}
	}
	return coordMap
}
