package marbles

// This will hold the state of the game in a way that makes it easy
// to manipulate. For this, I'm going to use a circular linked list.
// so I'll just create the functions for that.

type node struct {
	value int
	prev  *node
	next  *node
}

type CircularLinkedList struct {
	current *node
}

func NewCircularLinkedList() *CircularLinkedList {
	beginning := &node{value: 0}
	beginning.prev = beginning
	beginning.next = beginning
	return &CircularLinkedList{current: beginning}
}

// traverse the CLL clockwise
func (cll *CircularLinkedList) MoveCW(steps int) {
	for i := 0; i < steps; i++ {
		cll.current = cll.current.next
	}
}

func (cll *CircularLinkedList) MoveCCW(steps int) {
	for i := 0; i < steps; i++ {
		cll.current = cll.current.prev
	}
}

// insert new node at current state, which will insert
// AFTER the current state, so make sure you move CW once
func (cll *CircularLinkedList) Insert(value int) {
	newNode := &node{value: value}
	newNode.prev = cll.current
	newNode.next = cll.current.next
	cll.current.next.prev = newNode
	cll.current.next = newNode
	cll.current = newNode
}

// pop out of the CLL and return the value
func (cll *CircularLinkedList) Remove() int {
	value := cll.current.value
	cll.current.prev.next = cll.current.next
	cll.current.next.prev = cll.current.prev
	cll.current = cll.current.next

	return value
}
