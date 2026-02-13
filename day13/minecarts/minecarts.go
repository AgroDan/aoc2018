package minecarts

import "github.com/AgroDan/aocutils"

// This is the "next turn" enums
const (
	LEFT     = 0
	STRAIGHT = 1
	RIGHT    = 2
)

// The "direction" enums are defined
// in aocutils.Coord, so aocutils.N == 0,
// aocutils.E == 1, etc

type Cart struct {
	aocutils.Coord
	Direction int
	NextTurn  aocutils.Deque[int]
}

func NewCart(x, y int, direction int) *Cart {
	deque := aocutils.NewDeque[int]()
	for _, turn := range []int{LEFT, STRAIGHT, RIGHT} {
		deque.PushBack(turn)
	}

	return &Cart{
		Coord:     aocutils.Coord{X: x, Y: y},
		Direction: direction,
		NextTurn:  deque,
	}
}

func (c *Cart) chooseDirection() int {
	// This will return the next direction we're
	// going to turn, then pop it off the list and
	// add it to the end
	turn, success := c.NextTurn.PopFront()
	if !success {
		panic("Cart has no more turns to choose from")
	}

	c.NextTurn.PushBack(turn)
	return turn
}

func (c Cart) peekAhead() aocutils.Coord {
	// this will return the coordinate of the next position
	// in the carts path.

	//omg i wrote this already LOL
	return c.Peek(c.Direction)
}

func (c *Cart) turn(dir int) {
	switch c.Direction {
	case aocutils.N:
		switch dir {
		case LEFT:
			c.Direction = aocutils.W
		case RIGHT:
			c.Direction = aocutils.E
		}
	case aocutils.E:
		switch dir {
		case LEFT:
			c.Direction = aocutils.N
		case RIGHT:
			c.Direction = aocutils.S
		}
	case aocutils.S:
		switch dir {
		case LEFT:
			c.Direction = aocutils.E
		case RIGHT:
			c.Direction = aocutils.W
		}
	case aocutils.W:
		switch dir {
		case LEFT:
			c.Direction = aocutils.S
		case RIGHT:
			c.Direction = aocutils.N
		}
	}
}
