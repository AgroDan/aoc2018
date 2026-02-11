package fuelcells

import "fmt"

// Find the fuel cell's rack ID, which is its X coordinate plus 10.
// Begin with a power level of the rack ID times the Y coordinate.
// Increase the power level by the value of the grid serial number (your puzzle input).
// Set the power level to itself multiplied by the rack ID.
// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
// Subtract 5 from the power level.

// I'm not going to use my aocutils.coord object for this because that might be overkill

type Cell struct {
	X, Y int
}

func (c Cell) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func (c Cell) neighbors() []*Cell {
	var retval []*Cell
	for x := c.X - 1; x <= c.X+1; x++ {
		for y := c.Y - 1; y <= c.Y+1; y++ {
			if c.X == x && c.Y == y {
				continue
			}

			thisCell := Cell{
				X: x,
				Y: y,
			}
			retval = append(retval, &thisCell)
		}
	}
	return retval
}

func powerLevel(c Cell, serial int) int {
	// This finds the power level based on the equation above
	rackID := c.X + 10
	pl := rackID * c.Y
	pl += serial
	pl *= rackID
	if pl < 100 {
		pl = 0
	} else {
		pl /= 100
		pl %= 10
	}
	pl -= 5
	return pl
}

func totalPowerLevel(seed Cell, serial int) int {
	// This function gets the neighbors of the cell and sums up
	// their power levels so I can get a total power level for
	// each point in the 300x300 grid. MAKE SURE YOU DON'T CHOOSE
	// A BORDER POINT! It assumes the neighbors are all valid points
	var total int
	for _, neighbor := range seed.neighbors() {
		total += powerLevel(*neighbor, serial)
	}
	total += powerLevel(seed, serial)
	return total
}

func FindMaxPower(serial int) (Cell, int) {
	var maxPower int
	var maxCell Cell
	for x := 1; x < 300; x++ {
		for y := 1; y < 300; y++ {
			thisCell := Cell{
				X: x,
				Y: y,
			}
			thisPower := totalPowerLevel(thisCell, serial)
			if thisPower > maxPower {
				maxPower = thisPower
				maxCell = thisCell
			}
		}
	}
	// maybe a fix?
	maxCell.X--
	maxCell.Y--
	return maxCell, maxPower
}

// For part 2, I need to find the total power level for each cell and its neighbors,
// but I also need to find the total power level for each cell and its neighbors'
// neighbors, and so on up to 3x3, 4x4, etc. Time for some memoization. Maybe I can
// just create a map of every single cell with its power level.

func GeneratePowerLevelMap(serial int) [][][]int {
	// In this case, the first dimension is the X coordinate
	// the second dimension is the Y coordinate, and the third
	// dimension is the size of the square, [0] is the power level
	// of the cell itself, [1] is the power level of the cell and
	// its neighbors, [2] is the power level of the cell and its
	// neighbors and their neighbors, etc. Basically the layers
	// of the square.
	powerLevelMap := make([][][]int, 300)
	for x := 0; x < 300; x++ {
		powerLevelMap[x] = make([][]int, 300)
		for y := 0; y < 300; y++ {
			thisCell := Cell{
				X: x,
				Y: y,
			}
			powerLevelMap[x][y] = make([]int, 300)
			powerLevelMap[x][y][0] = powerLevel(thisCell, serial)
		}
	}

	// Now that the first part is done, I can fill in the rest of the layers.
	// for each layer, I'll just add the power level of the new cells to the
	// previous layer's power level. For example, for layer 1, I'll add the power
	// level of the neighbors to the power level of the cell itself. For layer 2,
	// I'll add the power level of the neighbors' neighbors to the power level of
	// the cell and its neighbors, etc. To kinda speed this up I'll stop growing
	// the layers once it hits a border
	for layer := 1; layer < 300; layer++ {
		for x := 0; x < 300; x++ {
			for y := 0; y < 300; y++ {
				if x+layer >= 300 || y+layer >= 300 {
					continue
				}
				powerLevelMap[x][y][layer] = powerLevelMap[x][y][layer-1] // we can find the area here!
				for i := 0; i <= layer; i++ {
					// This part will actually build the layer calculations.
					// We need to add the power level of the new cells to the
					// previous layer's power level. For example, for layer 1,
					// I'll add the power level of the neighbors to the power
					// level of the cell itself. For layer 2, I'll add the power
					// level of the neighbors' neighbors to the power level of the
					// cell and its neighbors, etc.
					powerLevelMap[x][y][layer] += powerLevelMap[x+i][y+layer][0]
					powerLevelMap[x][y][layer] += powerLevelMap[x+layer][y+i][0]
				}
			}
		}
	}
	return powerLevelMap
}

func FindBestPowerLevelPart2(serial int) (Cell, int, int) {
	powerLevelMap := GeneratePowerLevelMap(serial)
	var maxPower int
	var maxCell Cell
	var maxSize int
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			for size := 0; size < 300; size++ {
				thisPower := powerLevelMap[x][y][size]
				if thisPower > maxPower {
					maxPower = thisPower
					maxCell = Cell{X: x, Y: y}
					maxSize = size + 1 // because the size is actually the layer, so we need to add 1
					// to get the actual size of the square
				}
			}
		}
	}
	return maxCell, maxSize, maxPower
}
