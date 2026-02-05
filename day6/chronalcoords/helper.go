package chronalcoords

import (
	"fmt"
	"utils"
)

func (m *Ccmap) PrintMap() {
	for y := m.lY; y <= m.uY; y++ {
		for x := m.lX; x <= m.uX; x++ {
			knownCoord := false
			for _, coord := range m.coords {
				if coord.X == x && coord.Y == y {
					fmt.Printf("X")
					knownCoord = true
					break
				}
			}
			if !knownCoord {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (m *Ccmap) PrintEquallyFar() {
	for y := m.lY; y <= m.uY; y++ {
	inner:
		for x := m.lX; x <= m.uX; x++ {
			for _, coord := range m.coords {
				if coord.X == x && coord.Y == y {
					fmt.Printf("X")
					continue inner
				}
			}
			c := utils.Coord{X: x, Y: y}
			_, isUnique := m.FindClosestPoint(c)
			if !isUnique {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
