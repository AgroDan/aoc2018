package sand

import (
	"github.com/AgroDan/aocutils"
)

type ClayMap struct {
	aocutils.Runemap
	MinX, MaxX, MinY, MaxY int
	// waterSources           map[aocutils.Coord]WaterFall
}

func (cm *ClayMap) PrintOffset() {
	// This will print the map without so much extraneous data

	// the way I'm going to do this is creating an ephemeral
	// runemap then just printing that
	eph := aocutils.GenerateRunemap(cm.MaxX-cm.MinX+6, cm.MaxY-cm.MinY+1, '.')
	for y := cm.MinY; y <= cm.MaxY; y++ {
		for x := cm.MinX - 2; x <= cm.MaxX+2; x++ {
			getThisItem := aocutils.Coord{X: x, Y: y}
			setThisItem := aocutils.Coord{X: x - cm.MinX + 1, Y: y - cm.MinY}
			thisRune, err := cm.Get(getThisItem)
			if err != nil {
				thisRune = '.'
			}
			eph.Set(setThisItem, thisRune)
		}
	}
	eph.Print()
}
