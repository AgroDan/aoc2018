package sand

import (
	"slices"

	"github.com/AgroDan/aocutils"
)

func (c *ClayMap) waterInBounds(coord aocutils.Coord) bool {
	// Water can exist from Y=0 upward, not constrained by MinY
	return coord.Y > 0 && coord.Y <= c.MaxY
}

func (c *ClayMap) inBounds(coord aocutils.Coord) bool {
	return coord.X >= c.MinX && coord.X <= c.MaxX && coord.Y >= c.MinY && coord.Y <= c.MaxY
}

func (c *ClayMap) startInBounds(coord aocutils.Coord) bool {
	return coord.X >= c.MinX && coord.X <= c.MaxX && coord.Y <= c.MaxY
}

func (c *ClayMap) bfsSpread(seed aocutils.Coord) []aocutils.Coord {
	// this is BFS and will be weighted only to go
	// left and right of the point. If any drop off
	// points are found (meaning a place for water
	// to spill over) then those points will be
	// returned. Otherwise a 0-length slice will
	// be returned if contained.

	queue := aocutils.NewQueue[aocutils.Coord]()
	visited := make(map[aocutils.Coord]bool)
	dropoffs := []aocutils.Coord{}

	queue.Enqueue(seed)
	visited[seed] = true

	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()

		// Check left and right
		lr := []int{-1, 1}
		for _, dir := range lr {
			neighbor := aocutils.Coord{X: current.X + dir, Y: current.Y}

			// simple check, if we're either out of bounds or visited
			// previously, skip
			if !c.waterInBounds(neighbor) || visited[neighbor] {
				continue
			}

			// if we hit anything but a '.' or '|' then skip
			neighborVal, err := c.Get(neighbor)
			if err != nil {
				// if we're out of bounds, then it's a dropoff point
				// fmt.Printf("Found OOB dropoff at %v\n", neighbor)
				dropoffs = append(dropoffs, neighbor)
				continue
			}
			if neighborVal != '.' && neighborVal != '|' {
				continue
			}

			// if we can go down from this neighbor, then it's a dropoff point
			downNeighbor := aocutils.Coord{X: neighbor.X, Y: neighbor.Y + 1}
			downVal, err := c.Get(downNeighbor)
			if err != nil {
				// if we're out of bounds, then it's a dropoff point
				dropoffs = append(dropoffs, neighbor)
				continue
			}
			// need to check if we've defined this downfall already
			if downVal == '.' || downVal == '|' {
				dropoffs = append(dropoffs, neighbor)
				// fmt.Printf("Found dropoff at %v\n", aocutils.Coord{X: neighbor.X - 494, Y: neighbor.Y - 1})
				// fmt.Printf("Found dropoff at %v\n", neighbor)
				continue
			}
			visited[neighbor] = true
			queue.Enqueue(neighbor)
		}
	}

	return dropoffs
}

func (c *ClayMap) fill(seed aocutils.Coord, downfalls []aocutils.Coord) {
	// like bfsspread, this will fill left and right of the seed.
	// this will NOT check to see if there are dropoffs. Instead it
	// will just fill everything it can.
	queue := aocutils.NewQueue[aocutils.Coord]()
	visited := make(map[aocutils.Coord]bool)

	var fillRune rune
	if len(downfalls) == 0 {
		fillRune = '~'
	} else {
		fillRune = '|'
	}

	queue.Enqueue(seed)
	visited[seed] = true

	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()
		c.Set(current, fillRune)

		if slices.Contains(downfalls, current) {
			continue
		}

		for _, dir := range []int{-1, 1} {
			neighbor := aocutils.Coord{X: current.X + dir, Y: current.Y}

			if !c.waterInBounds(neighbor) || visited[neighbor] {
				continue
			}

			neighborVal, err := c.Get(neighbor)
			if err != nil {
				// fmt.Printf("Error getting neighbor %v: %v\n", neighbor, err)
				continue
			}
			if neighborVal == '#' {
				continue
			}

			visited[neighbor] = true
			queue.Enqueue(neighbor)
		}
	}
}

func (c *ClayMap) Tick(activeWater []aocutils.Coord) []aocutils.Coord {
	// this will be one tick of the simulation. It will return a slice of
	// new seeds to be processed on the next tick. If the slice is empty,
	// then the simulation is done.
	// There are 4 possible runes to deal with:
	// . = empty
	// # = clay
	// | = water falling
	// ~ = water settled
	//
	// So for each "tick", I will ask the following in this order
	// for each "Waterfalling" point:
	// 1. Can I fall? - cell below is air, move and stay waterfalling
	// 2. Can I spread? - cell below is wall or water settled. Look for cell with air below it
	// 3. Am I contained? - use bfsspread to find if there are any places for water to fall.
	var retval []aocutils.Coord

	for _, seed := range activeWater {
		// First, look down. Can I fall?
		// Allow X out of bounds (water escaping the map)
		if seed.Y > c.MaxY {
			continue
		}
		downCoord := seed.Peek(aocutils.S)
		peekDown, err := c.Get(downCoord)
		// Water falls if: cell below is empty OR cell below is out of bounds (past MaxY)
		if (peekDown == '.' || peekDown == '|') || (err != nil &&
			(downCoord.Y > c.MaxY || downCoord.X > c.MaxX || downCoord.X < c.MinX)) {
			// if err != nil {
			// 	fmt.Printf("%v\n", err)
			// }
			// Note, adding | for peeking down due to edge cases where
			// water can fall into a pool that's already been filled.
			c.Set(seed, '|')
			// Only continue falling if the cell is in bounds
			// also if we're not falling into running water again
			// save us some ticks
			if c.waterInBounds(downCoord) && peekDown != '|' {
				retval = append(retval, downCoord)
			}
			continue
		}

		// otherwise, can I spread?
		if peekDown == '#' || peekDown == '~' {
			downfalls := c.bfsSpread(seed)
			c.fill(seed, downfalls)
			retval = append(retval, downfalls...)

			if len(downfalls) == 0 {
				// We're filling, so backtrack
				retval = append(retval, seed.Peek(aocutils.N))
			}
		}
	}

	return retval
}

func (c *ClayMap) CountWater() int {
	count := 0
	for y := c.MinY; y <= c.MaxY; y++ {
		for x := 0; x <= c.Width(); x++ {
			r, _ := c.Get(aocutils.Coord{X: x, Y: y})
			if r == '~' || r == '|' {
				count++
			}
		}
	}
	return count
}

func (c *ClayMap) CountSettledWater() int {
	count := 0
	for y := c.MinY; y <= c.MaxY; y++ {
		for x := 0; x <= c.Width(); x++ {
			r, _ := c.Get(aocutils.Coord{X: x, Y: y})
			if r == '~' {
				count++
			}
		}
	}
	return count
}
