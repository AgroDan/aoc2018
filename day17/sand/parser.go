package sand

import (
	"fmt"

	"github.com/AgroDan/aocutils"
)

// This will just parse each line and return
// the coordinates of the clay as a slice of points.

func registerLine(line string) []aocutils.Coord {
	var points []aocutils.Coord
	var fixed, from, to int
	var fixedChar, rangeChar byte
	fmt.Sscanf(line, "%c=%d, %c=%d..%d", &fixedChar, &fixed, &rangeChar, &from, &to)

	for i := from; i <= to; i++ {
		switch fixedChar {
		case 'x':
			points = append(points, aocutils.Coord{X: fixed, Y: i})
		case 'y':
			points = append(points, aocutils.Coord{X: i, Y: fixed})
		}
	}

	return points
}

func parseInput(input []string) (int, int, int, int, [][]aocutils.Coord) {
	// This will return the minimum x and y coordinates, the maximum x
	// and y coordinates, and a slice of slices of points representing the
	// clay. At this point I can create a runemap and an offset so we can
	// print the map nicer I guess.
	var clay [][]aocutils.Coord
	var minX, minY, maxX, maxY int
	minX = 9999
	minY = 9999
	maxX = -9999
	maxY = -9999

	for _, line := range input {
		points := registerLine(line)
		clay = append(clay, points)

		for _, point := range points {
			if point.X < minX {
				minX = point.X
			}
			if point.X > maxX {
				maxX = point.X
			}
			if point.Y < minY {
				minY = point.Y
			}
			if point.Y > maxY {
				maxY = point.Y
			}
		}
	}

	return minX, minY, maxX, maxY, clay
}

func GenerateMap(input []string) *ClayMap {
	// Does the heavy lifting.
	// minX, _, maxX, maxY, clay := parseInput(input)
	minX, minY, maxX, maxY, clay := parseInput(input)

	// var waterStart = make(map[aocutils.Coord]WaterFall)
	// waterStart[source] = WaterFall{stream: []aocutils.Coord{source}, ended: false}

	// on second thought, I'll generate the map without
	// any offsets just to keep things simple. Printing the
	// map will have the proper offsets I guess.
	cm := &ClayMap{
		Runemap: aocutils.GenerateRunemap(maxX+3, maxY+1, '.'),
		MinX:    minX,
		MaxX:    maxX,
		MinY:    minY,
		MaxY:    maxY,
		// waterSources: waterStart,
	}

	for _, points := range clay {
		for _, point := range points {
			cm.Set(point, '#')

		}
	}

	return cm
}
