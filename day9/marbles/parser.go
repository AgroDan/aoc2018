package marbles

import "fmt"

func Parse(input string) *Game {
	var players, marbles int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &marbles)

	return NewGame(players, marbles)
}
