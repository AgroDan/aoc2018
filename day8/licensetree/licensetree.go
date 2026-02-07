package licensetree

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AgroDan/aocutils"
)

// First, I need to create the tree structure,
// then I'll import the newly-created deque package
// that I wrote in the aocutils repository (trying to
// break away from using a local utils package so I can
// make changes to a central area rather than having to
// copy paste changes across multiple repositories).

type Node struct {
	Children []*Node
	Metadata []int
}

// Not necessary since I'll just be adding
// things regardless of their size of children, metadata, etc
func NewNode(header, metadata int) *Node {
	return &Node{
		Children: make([]*Node, header),
		Metadata: make([]int, metadata),
	}
}

// Now I'll create a function to create a node
// based on the current deque state
func BuildNode(dq *aocutils.Deque[int]) *Node {
	if dq.Size() < 2 {
		panic("Not enough data to build node header")
	}

	header, success := dq.PopFront()
	if !success {
		panic("Error popping header from deque")
	}

	metadataCount, success := dq.PopFront()
	if !success {
		panic("Error popping metadata count from deque")
	}

	node := NewNode(header, metadataCount)

	// Now add the metadata entries
	for i := 0; i < metadataCount; i++ {
		if dq.IsEmpty() {
			panic("Not enough data to pop metadata")
		}
		metadataEntry, success := dq.PopBack()
		if !success {
			panic("Error popping metadata entry from deque")
		}
		node.Metadata[i] = metadataEntry
	}

	return node
}

// Now a function to arbitrarily add a child node to a parent node
func (n *Node) AddChild(child *Node) error {
	for i := range n.Children {
		if n.Children[i] == nil {
			n.Children[i] = child
			return nil
		}
	}
	return fmt.Errorf("No empty child slot available")
}

// After some googling, I discovered that this is best handled with recursion,
// and NOT a deque like I first thought. So let's work on this.

func parse(data []int, pos int) (*Node, int) {
	// Time to recurse like a mofo
	numChildren := data[pos]
	numMetadata := data[pos+1]
	pos += 2

	node := Node{
		Children: make([]*Node, 0), // setting to zero so I can just append
		Metadata: make([]int, 0),
	}

	for i := 0; i < numChildren; i++ {
		thischild, newPos := parse(data, pos)
		node.Children = append(node.Children, thischild)
		pos = newPos
	}

	for i := 0; i < numMetadata; i++ {
		node.Metadata = append(node.Metadata, data[pos])
		pos++
	}

	return &node, pos
}

func BuildTree(line string) *Node {
	stringNum := strings.Split(line, " ")
	var license []int
	for i := range stringNum {
		num, err := strconv.Atoi(stringNum[i])
		if err != nil {
			panic("Error converting string to int")
		}
		license = append(license, num)
	}

	root, _ := parse(license, 0)
	return root
}

func (n *Node) SumMetadata() int {
	total := 0
	for i := range n.Metadata {
		total += n.Metadata[i]
	}
	for i := range n.Children {
		if n.Children[i] != nil {
			total += n.Children[i].SumMetadata()
		}
	}
	return total
}

// And now even MORE recursion for part 2!
// Basically to work backwards, I'm going to
// assign a function to a node, and if it has
// no children, then it will just sum its metadata.
// If it has children, then it will refer to each
// child node's value based on the metadata entries
// and sum those.

func (n *Node) Value() int {
	if len(n.Children) == 0 {
		return n.SumMetadata()
	}

	// Otherwise, time to recurse.
	total := 0
	for i := range n.Metadata {
		childIndex := n.Metadata[i] - 1
		// eric wastl is a perl programmer and thinks
		// we should index at 1 instead of 0 but we
		// can forgive him I GUESS
		if childIndex >= 0 && childIndex < len(n.Children) {
			total += n.Children[childIndex].Value()
		}
	}
	return total
}
