package sort

import "fmt"

// Node represents a single node in the list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList holds the head of the list
type LinkedList struct {
	Head *Node
}

// Insert adds a new node at the end
func (l *LinkedList) Insert(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

// Delete removes the first node with the given value
func (l *LinkedList) Delete(value int) bool {
	if l.Head == nil {
		return false
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return true
	}

	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.Value == value {
			prev.Next = curr.Next
			return true
		}
		prev = curr
		curr = curr.Next
	}

	return false
}

// Search checks if a value exists in the list
func (l *LinkedList) Search(value int) bool {
	curr := l.Head
	for curr != nil {
		if curr.Value == value {
			return true
		}
		curr = curr.Next
	}
	return false
}

// Print displays the list elements
func (l *LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// Length returns the number of nodes
func (l *LinkedList) Length() int {
	count := 0
	curr := l.Head
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
