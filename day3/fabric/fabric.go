package fabric

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Fabric struct {
	ID            int
	Width, Height int
	Pos           utils.Coord
}

func NewFabric(line string) *Fabric {
	// This will ingest the line looking like:
	// #123 @ 3,2: 5x4
	// which means: ID of 123
	// positioned at X: 3, Y: 2 (remember building down)
	// 5 inches wide, 4 inches tall
	lineParts := strings.Split(line, " ")
	f := &Fabric{}
	f.ID, _ = strconv.Atoi(strings.TrimPrefix(lineParts[0], "#"))

	// Now position, let's try sscanf this time
	var x, y int
	fmt.Sscanf(lineParts[2], "%d,%d:", &x, &y)
	f.Pos = utils.Coord{X: x, Y: y}

	var wid, hei int
	fmt.Sscanf(lineParts[3], "%dx%d", &wid, &hei)
	f.Width, f.Height = wid, hei
	return f
}

func (f Fabric) String() string {
	return fmt.Sprintf("ID: %d, Position: X: %d, Y: %d, Width: %d, Height: %d", f.ID, f.Pos.X, f.Pos.Y, f.Width, f.Height)
}

func (f Fabric) Overlaps(other Fabric) bool {
	// Check if this fabric overlaps with another fabric
	// Using Axis-Aligned Bounding Box (AABB) collision detection
	// This is a good explanation here: https://www.youtube.com/watch?v=YbWkPcdo8fk
	if f.Pos.X < other.Pos.X+other.Width &&
		f.Pos.X+f.Width > other.Pos.X &&
		f.Pos.Y < other.Pos.Y+other.Height &&
		f.Pos.Y+f.Height > other.Pos.Y {
		return true
	}
	return false
}

// Now get the calculation of exactly how much overlap. Had to google this
// but it seems pretty simple

// Note that it should be determined ahead of time if there actually is some
// overlap, otherwise we'll get some weird numbers
func (f Fabric) HowMuchOverlap(other Fabric) (int, int) {
	// Returns Width, Height
	var overlapX, overlapY int

	overlapX = utils.Min(f.Pos.X+f.Width, other.Pos.X+other.Width) - utils.Max(f.Pos.X, other.Pos.X)
	overlapY = utils.Min(f.Pos.Y+f.Height, other.Pos.Y+other.Height) - utils.Max(f.Pos.Y, other.Pos.Y)
	return overlapX, overlapY
}
