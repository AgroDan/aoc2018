package chronalcoords

import "utils"

// This will work alongside with the main mapper to determine
// a list of coordinates that we can test the manhattan distance
// from.

func (m *Ccmap) FindClosestPoint(c utils.Coord) (utils.Coord, bool) {
	minDist := -1
	var closestPoints []utils.Coord
	for _, point := range m.coords {
		dist := utils.ManhattanDistance(c, point)
		if minDist == -1 || dist < minDist {
			minDist = dist
			closestPoints = []utils.Coord{point}
		} else if dist == minDist {
			closestPoints = append(closestPoints, point)
		}
	}
	if len(closestPoints) == 1 {
		return closestPoints[0], true
	}
	// tie, no one wins
	return utils.Coord{X: -1, Y: -1}, false
}

// This function will create a box of the boundary coordinates.
func (m *Ccmap) generateBoundaryPoints() []utils.Coord {
	var boundaryPoints []utils.Coord
	// top and bottom rows
	for x := m.lX; x <= m.uX; x++ {
		boundaryPoints = append(boundaryPoints, utils.Coord{X: x, Y: m.lY})
		boundaryPoints = append(boundaryPoints, utils.Coord{X: x, Y: m.uY})
	}
	// left and right columns
	for y := m.lY + 1; y < m.uY; y++ {
		boundaryPoints = append(boundaryPoints, utils.Coord{X: m.lX, Y: y})
		boundaryPoints = append(boundaryPoints, utils.Coord{X: m.uX, Y: y})
	}
	return boundaryPoints
}

// This function will find points that have "infinte area", basically
// any point that is closest to a boundary point.
func (m *Ccmap) FindInfiniteAreaPoints() map[utils.Coord]bool {
	infinitePoints := make(map[utils.Coord]bool)
	boundaryPoints := m.generateBoundaryPoints()
	for _, bPoint := range boundaryPoints {
		closestPoint, _ := m.FindClosestPoint(bPoint)
		if closestPoint.X != -1 && closestPoint.Y != -1 {
			infinitePoints[closestPoint] = true
		}
	}
	return infinitePoints
}

// And finally, this will check every single point in the boundary box
// to see how many points are closest to each coordinate. Eventually
// I'll sort these out to find the largest area, then confirm that the
// largest one isn't infinite.
func (m *Ccmap) CalculateClosestAreas() map[utils.Coord]int {
	// This var will hold each coordinate and how many points
	// within the boundary are closest to it
	areaCounts := make(map[utils.Coord]int)

	for x := m.lX; x <= m.uX; x++ {
		for y := m.lY; y <= m.uY; y++ {
			currentPoint := utils.Coord{X: x, Y: y}

			// remember, some points may tie so ignore them
			closestPoint, isUnique := m.FindClosestPoint(currentPoint)
			if isUnique {
				areaCounts[closestPoint]++
			}
		}
	}
	return areaCounts
}

// And now for part 2, we'll look through every single point in the map
// and see if the sum of the manhattan distances to all coordinates
// is less than a given threshold.
func (m *Ccmap) CalculateRegionSize(maxDistance int) int {
	regionSize := 0

	for x := m.lX; x <= m.uX; x++ {
		for y := m.lY; y <= m.uY; y++ {
			currentPoint := utils.Coord{X: x, Y: y}

			totalDistance := 0
			for _, coord := range m.coords {
				totalDistance += utils.ManhattanDistance(currentPoint, coord)
			}

			if totalDistance < maxDistance {
				regionSize++
			}
		}
	}

	return regionSize
}
