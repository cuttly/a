package main

// Write a function ListSize that returns the number of elements in a linked list l.

// COMMENT OUT THE BELOW List and NodeL struct if needed
type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}

// COMMENT OUT THE ABOVE List and NodeL struct if needed

func ListSize(l *List) int {
	if l.Head == nil { // if the list is empty
		return 0
	}
	count := 1
	for l.Head.Next != nil {
		count++              // otherwise, count this existing element
		l.Head = l.Head.Next // and move to the next element
	}
	return count
}
