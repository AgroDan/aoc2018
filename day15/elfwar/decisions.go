package elfwar

import "github.com/AgroDan/aocutils"

func readingOrder(coords *[]aocutils.Coord) {
	// Sorts a set of coordinates in "reading order", and this
	// just modifies the slice in place
	for i := 0; i < len(*coords)-1; i++ {
		for j := 0; j < len(*coords)-i-1; j++ {
			if (*coords)[j].Y > (*coords)[j+1].Y || ((*coords)[j].Y == (*coords)[j+1].Y && (*coords)[j].X > (*coords)[j+1].X) {
				(*coords)[j], (*coords)[j+1] = (*coords)[j+1], (*coords)[j]
			}
		}
	}
}
