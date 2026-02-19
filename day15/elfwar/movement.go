package elfwar

import (
	"container/heap"
	"math"

	"github.com/AgroDan/aocutils"
)

// something tells me this may need its own file so it doesn't clutter everything up

// func (b *Battlefield) FindShortestPath(start, end aocutils.Coord) ([]aocutils.Coord, bool) {
// I think the best approach here is A* because it's the least computationally
// expensive. BFS might be overkill because it's just gonna fill everything in
// and I did A* a while back, so let's emulate that.

// The basic idea here is that _this_ will give me the shortest path from start to end.

// specifically I'm copying (mostly) the code I used in AOC2024, day 18: Ram Run

type Node struct {
	aocutils.Coord
	G, H   float64
	Parent *Node
}

func (n *Node) F() float64 {
	return n.G + n.H
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F() < pq[j].F()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func heuristic(a, b aocutils.Coord) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

// This is necessary specific to the battlefield above to determine
// if the point contains a valid space or not. As far as I'm concerned,
// the only valid spaces are the empty space, '.'
// NOTE: MAY NEED TO ACCOUNT FOR OTHER UNITS AS WELL
func isValid(point aocutils.Coord, b *Battlefield) bool {
	// also need to check bounds
	if point.X < 0 || point.Y < 0 || point.X >= b.Width() || point.Y >= b.Height() {
		return false
	}

	// Is a unit present?
	if b.IsUnitPresentAndAlive(point) {
		return false
	}

	checkRune, err := b.Get(point)
	if err != nil {
		// cover my bases
		return false
	}

	return checkRune == '.'
}

// helper for A*
func reconstructPath(node *Node) []aocutils.Coord {
	var path []aocutils.Coord
	for node != nil {
		path = append([]aocutils.Coord{node.Coord}, path...)
		node = node.Parent
	}
	return path
}

// here is where the magic happens
func aStar(start, end aocutils.Coord, b *Battlefield) ([]aocutils.Coord, float64) {
	openSet := &PriorityQueue{}
	closedSet := make(map[aocutils.Coord]struct{})
	var beenThere struct{}

	heap.Init(openSet)
	heap.Push(openSet, &Node{
		Coord: start,
		G:     0,
		// H:      heuristic(start, end),
		H:      float64(bfsHeuristic(start, end, b)),
		Parent: nil,
	})

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)

		if current.Coord == end {
			return reconstructPath(current), current.G
		}

		closedSet[current.Coord] = beenThere

		// check neighbors, note ONLY up/down/left/right
		directions := current.Coord.TrueAllAvailable()

		for _, neighbor := range directions {
			if !isValid(neighbor, b) {
				continue
			}

			if _, exists := closedSet[neighbor]; exists {
				continue
			}

			gScore := current.G + 1
			hScore := float64(bfsHeuristic(neighbor, end, b))

			// check if this is a new node
			found := false
			for _, node := range *openSet {
				if node.Coord == neighbor {
					found = true
					if gScore < node.G {
						node.G = gScore
						node.Parent = current
					}
				}
			}

			if !found {
				heap.Push(openSet, &Node{
					Coord:  neighbor,
					G:      gScore,
					H:      hScore,
					Parent: current,
				})
			}
		}
	}

	return nil, 0
}

// Here's a wild idea. I'm going to use a BFS function to determine how many
// steps it would take to get to each point. This will be useful for determining
// which point to move to.
func bfsHeuristic(start, end aocutils.Coord, b *Battlefield) int {
	visited := make(map[aocutils.Coord]struct{})
	var beenThere struct{}

	type step struct {
		Coord aocutils.Coord
		Steps int
	}

	queue := aocutils.NewQueue[step]()
	queue.Enqueue(step{Coord: start, Steps: 0})
	visited[start] = beenThere

	steps := 0
	lowestSteps := math.MaxInt64
	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()
		coord := current.Coord
		steps = current.Steps
		if steps > lowestSteps {
			continue
		}
		if coord == end {
			if steps < lowestSteps {
				lowestSteps = steps
			}
		}

		directions := coord.TrueAllAvailable()
		for _, neighbor := range directions {
			if !isValid(neighbor, b) {
				continue
			}

			if _, exists := visited[neighbor]; exists {
				continue
			}

			visited[neighbor] = beenThere
			thisStep := step{Coord: neighbor, Steps: steps + 1}
			queue.Enqueue(thisStep)
		}
		steps++
	}

	if lowestSteps == math.MaxInt64 {
		return -1
	}
	return lowestSteps
}
