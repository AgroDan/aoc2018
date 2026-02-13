package minecarts

import (
	"fmt"

	"github.com/AgroDan/aocutils"
)

// This will be the track struct,
// handling all of the carts and
// allowing them to move

type Track struct {
	Carts []*Cart
	aocutils.Runemap
}

func NewTrack(input []string) *Track {
	track := &Track{
		Carts:   []*Cart{},
		Runemap: aocutils.NewRunemap(input),
	}

	leftFacing := track.FindAll('<')
	for _, c := range leftFacing {
		track.Carts = append(track.Carts, NewCart(c.X, c.Y, aocutils.W))
		track.Set(c, '-') // we know the cart is on a horizontal track
		// and the carts location so let's just change the track to the
		// proper track piece
	}

	rightFacing := track.FindAll('>')
	for _, c := range rightFacing {
		track.Carts = append(track.Carts, NewCart(c.X, c.Y, aocutils.E))
		track.Set(c, '-')
	}

	upFacing := track.FindAll('^')
	for _, c := range upFacing {
		track.Carts = append(track.Carts, NewCart(c.X, c.Y, aocutils.N))
		track.Set(c, '|')
	}

	downFacing := track.FindAll('v')
	for _, c := range downFacing {
		track.Carts = append(track.Carts, NewCart(c.X, c.Y, aocutils.S))
		track.Set(c, '|')
	}

	return track
}

func (t *Track) Tick() aocutils.Coord {
	// This moves all the carts one tick, in the correct order, and handles
	// all the turns and intersection decisions etc. To get it in the right
	// order, first I need to sweep the map from the top left, moving right
	// until the end of the line, then down to the next line, etc. If I find
	// a cart in that position, push it onto the queue.

	orderQueue := aocutils.NewQueue[*Cart]()
	for y := 0; y < t.Height(); y++ {
		for x := 0; x < t.Width(); x++ {
			for _, cart := range t.Carts {
				if cart.X == x && cart.Y == y {
					orderQueue.Enqueue(cart)
				}
			}
		}
	}

	// cartPositions := make(map[aocutils.Coord]bool)
	for !orderQueue.IsEmpty() {
		cart, _ := orderQueue.Dequeue()
		// first we need to peek ahead to see where the cart is going to move
		nextCoord := cart.peekAhead()

		if t.isSpaceOccupied(nextCoord) {
			// collision!
			return nextCoord
		}

		// move the cart
		cart.X = nextCoord.X
		cart.Y = nextCoord.Y

		// now check the track piece to decide if we need to turn
		trackPiece, _ := t.Get(nextCoord)
		switch trackPiece {
		case '/':
			switch cart.Direction {
			case aocutils.N:
				cart.Direction = aocutils.E
			case aocutils.S:
				cart.Direction = aocutils.W
			case aocutils.E:
				cart.Direction = aocutils.N
			case aocutils.W:
				cart.Direction = aocutils.S
			}
		case '\\': // escape backslashes lol
			switch cart.Direction {
			case aocutils.N:
				cart.Direction = aocutils.W
			case aocutils.S:
				cart.Direction = aocutils.E
			case aocutils.E:
				cart.Direction = aocutils.S
			case aocutils.W:
				cart.Direction = aocutils.N
			}
		case '+':
			cart.turn(cart.chooseDirection())
		}
	}

	return aocutils.Coord{X: -1, Y: -1} // no collision
}

func PrintTrack(track *Track) {
	// this is just a helper function to print the track with the carts on it
	// for debugging purposes
	for y := 0; y < track.Height(); y++ {
		for x := 0; x < track.Width(); x++ {
			cartHere := false
			for _, cart := range track.Carts {
				if cart.X == x && cart.Y == y {
					cartHere = true
					switch cart.Direction {
					case aocutils.N:
						fmt.Print("^")
					case aocutils.S:
						fmt.Print("v")
					case aocutils.E:
						fmt.Print(">")
					case aocutils.W:
						fmt.Print("<")
					}
				}
			}
			if !cartHere {
				r, _ := track.Get(aocutils.Coord{X: x, Y: y})
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()
	}
}

func (t *Track) isSpaceOccupied(coord aocutils.Coord) bool {
	for _, cart := range t.Carts {
		if cart.X == coord.X && cart.Y == coord.Y {
			return true
		}
	}
	return false
}

func (t *Track) TickPartTwo() aocutils.Coord {
	// This is the same as Tick, but instead of returning
	// the first collision, will remove two carts after they
	// collide. Then, once there is only one cart left, return its
	// position

	// NOTE: Make sure you re-load the track before running this!

	orderQueue := aocutils.NewQueue[*Cart]()
	for y := 0; y < t.Height(); y++ {
		for x := 0; x < t.Width(); x++ {
			for _, cart := range t.Carts {
				if cart.X == x && cart.Y == y {
					orderQueue.Enqueue(cart)
				}
			}
		}
	}

	for !orderQueue.IsEmpty() {
		cart, _ := orderQueue.Dequeue()
		// first we need to peek ahead to see where the cart is going to move
		nextCoord := cart.peekAhead()

		if t.isSpaceOccupied(nextCoord) {
			// collision!
			t.removeCartAt(nextCoord)
			// and the moving cart
			t.removeCartAt(cart.Coord)
			continue
		}

		// move the cart
		cart.X = nextCoord.X
		cart.Y = nextCoord.Y

		// now check the track piece to decide if we need to turn
		trackPiece, _ := t.Get(nextCoord)
		switch trackPiece {
		case '/':
			switch cart.Direction {
			case aocutils.N:
				cart.Direction = aocutils.E
			case aocutils.S:
				cart.Direction = aocutils.W
			case aocutils.E:
				cart.Direction = aocutils.N
			case aocutils.W:
				cart.Direction = aocutils.S
			}
		case '\\': // escape backslashes lol
			switch cart.Direction {
			case aocutils.N:
				cart.Direction = aocutils.W
			case aocutils.S:
				cart.Direction = aocutils.E
			case aocutils.E:
				cart.Direction = aocutils.S
			case aocutils.W:
				cart.Direction = aocutils.N
			}
		case '+':
			cart.turn(cart.chooseDirection())
		}
	}

	if len(t.Carts) == 1 {
		return t.Carts[0].Coord
	}

	return aocutils.Coord{X: -1, Y: -1} // no collision
}

func (t *Track) removeCartAt(coord aocutils.Coord) {
	for i, cart := range t.Carts {
		if cart.X == coord.X && cart.Y == coord.Y {
			// i think this is clever
			t.Carts = append(t.Carts[:i], t.Carts[i+1:]...)
			return
		}
	}
}
