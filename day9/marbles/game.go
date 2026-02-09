package marbles

// This will play the actual game.

type Game struct {
	players int
	Marbles int
	scores  []int
	board   *CircularLinkedList
}

func NewGame(players, marbles int) *Game {
	return &Game{
		players: players,
		Marbles: marbles,
		scores:  make([]int, players),
		board:   NewCircularLinkedList(),
	}
}

func (g *Game) Play() {
	for i := 1; i <= g.Marbles; i++ {
		if i%23 == 0 {
			// rules stipulate that we need to move 7 marbles CCW
			// and then pop the value at that position, which will be added to the score
			g.board.MoveCCW(7)
			g.scores[i%g.players] += i + g.board.Remove()
		} else {
			g.board.MoveCW(1)
			g.board.Insert(i)
		}
	}
}

func (g *Game) HighScore() int {
	high := 0
	for _, score := range g.scores {
		if score > high {
			high = score
		}
	}
	return high
}
