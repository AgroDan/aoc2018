package beams

import (
	"fmt"

	"github.com/AgroDan/aocutils"
)

// This will define a beam as having a given coordinate
// and a direction. I'll have functions which will update
// the coordinates based on the direction per step, and
// then I'll have to figure out how the hell I'm going to
// display them when _all_ the beams are adjacent. I'll
// have to figure that out later.

type beam struct {
	aocutils.Coord
	velocityX, velocityY int
}

func newBeam(line string) *beam {
	var b beam
	fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &b.X, &b.Y, &b.velocityX, &b.velocityY)
	return &b
}

func (b *beam) step() {
	b.X += b.velocityX
	b.Y += b.velocityY
}

type BeamCollection []*beam

func NewBeamCollection(lines []string) BeamCollection {
	bc := make(BeamCollection, len(lines))
	for i, line := range lines {
		bc[i] = newBeam(line)
	}
	return bc
}

func (bc BeamCollection) CheckAdjacency() bool {
	// This is going to be a little similar to that tree bathroom robot challenge
	// in the 2024 AoC in that I'll have to check all of the coordinates in this
	// collection and see if they're all adjacent. I can do this by checking the
	// distance between each pair of beams and seeing if it's less than or equal to 1.
	for i := range bc {
		isAdjacent := false
		for j := range bc {
			if i == j {
				continue
			}
			if aocutils.Abs(bc[i].X-bc[j].X) <= 1 && aocutils.Abs(bc[i].Y-bc[j].Y) <= 1 {
				isAdjacent = true
				break
			}
		}
		if !isAdjacent {
			return false
		}
	}
	return true
}

func (bc BeamCollection) Step() {
	for _, b := range bc {
		b.step()
	}
}

func (bc BeamCollection) Display() {
	// This is going to be a little tricky. I'll have to find the min and max coordinates
	// to determine the size of the grid, and then I'll have to create a grid and fill it
	// with the beams. I'll have to be careful with the coordinates since they can be negative.
	minX, maxX := bc[0].X, bc[0].X
	minY, maxY := bc[0].Y, bc[0].Y
	for _, b := range bc {
		if b.X < minX {
			minX = b.X
		}
		if b.X > maxX {
			maxX = b.X
		}
		if b.Y < minY {
			minY = b.Y
		}
		if b.Y > maxY {
			maxY = b.Y
		}
	}

	// Time for a runemap. Just for the sake of displaying this, I'll need to create an offset
	// first and use that to create a grid of the appropriate size.
	offsetX := -minX
	offsetY := -minY

	// Don't make fun of me, this is so I can retrofit this
	var runeMapString []string
	for y := minY; y <= maxY; y++ {
		row := ""
		for x := minX; x <= maxX; x++ {
			row += " "
		}
		runeMapString = append(runeMapString, row)
	}

	// OH THIS IS SO GHETTO but whatever I get to use a runemap
	runeMap := aocutils.NewRunemap(runeMapString)
	for _, b := range bc {
		setCoord := aocutils.Coord{X: b.X + offsetX, Y: b.Y + offsetY}
		runeMap.Set(setCoord, '#')
	}

	runeMap.Print()
}
