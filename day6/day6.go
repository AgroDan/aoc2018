package main

import (
	"day6/chronalcoords"
	"flag"
	"fmt"
	"time"
	"utils"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")
	// any additional flags add here
	threshold := flag.Int("t", 10000, "Threshold for part 2 region size")

	flag.Parse()

	// Choose based on the challenge...
	// individual lines:
	lines, err := utils.GetFileLines(*filePtr)
	if err != nil {
		fmt.Println("Fatal:", err)
	}

	// giant text blob:
	// challengeText, err := utils.GetTextBlob(*filePtr)
	// if err != nil {
	//     fmt.Println("Fatal:", err)
	// }

	// Insert code here
	mapper := chronalcoords.CreateMapper(lines)
	// fmt.Println("Possible coordinates to check:", mapper.GetPossibleCheckCoords())
	infinitePoints := mapper.FindInfiniteAreaPoints()
	areas := mapper.CalculateClosestAreas()

	// testCoord := utils.Coord{X: 5, Y: 2}
	// closestPoint, isUnique := mapper.FindClosestPoint(testCoord)
	// if isUnique {
	// 	fmt.Printf("Closest point to %v is %v\n", testCoord, closestPoint)
	// } else {
	// 	fmt.Printf("No unique closest point to %v\n", testCoord)
	// }

	largestArea := 0
	// topCoord := utils.Coord{}
	for coord, area := range areas {
		if _, isInfinite := infinitePoints[coord]; !isInfinite {
			if area > largestArea {
				largestArea = area
				// topCoord = coord
			}
		}
	}
	fmt.Println("Part 1 area size:", largestArea)
	// fmt.Println("Coordinate with largest area:", topCoord)
	// fmt.Println("Infinite area coordinates:", infinitePoints)
	// fmt.Println("All area sizes:", areas)

	// mapper.PrintMap()
	// fmt.Printf("\n")
	// mapper.PrintEquallyFar()
	regionSize := mapper.CalculateRegionSize(*threshold)
	fmt.Println("Part 2 region size:", regionSize)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
