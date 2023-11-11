package dlinklist

import "fmt"

// Node represents a node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	Head *Node
}

// NewLinkedList initializes and returns a new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds a new node with the given Data to the end of the linked list
func (ll *LinkedList) Append(Data int) {
	newNode := &Node{Data: Data, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Display prints the elements of the linked list
func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

// Search searches for a given value in the linked list
func (ll *LinkedList) Search(value int) *Node {
	current := ll.Head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// Delete removes the first occurrence of a value from the linked list
func (ll *LinkedList) Delete(value int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Data == value {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil && current.Next.Data != value {
		current = current.Next
	}

	if current.Next == nil {
		return
	}

	current.Next = current.Next.Next
}

// InsertAfter inserts a new node with the given Data after a specific node
func (ll *LinkedList) InsertAfter(prevNode *Node, Data int) {
	if prevNode == nil {
		fmt.Println("Previous node cannot be nil")
		return
	}

	newNode := &Node{Data: Data, Next: prevNode.Next}
	prevNode.Next = newNode
}

// Reverse reverses the linked list in-place
func (ll *LinkedList) Reverse() {
	var prev *Node = nil
	current := ll.Head
	var Next *Node = nil

	for current != nil {
		Next = current.Next
		current.Next = prev
		prev = current
		current = Next
	}

	ll.Head = prev
}

// Sort performs a simple bubble sort on the linked list
func (ll *LinkedList) Sort() {
	if ll.Head == nil {
		return
	}

	swapped := true
	for swapped {
		swapped = false
		current := ll.Head
		for current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				swapped = true
			}
			current = current.Next
		}
	}
}

// GetHead returns the Head of the linked list
func (list *LinkedList) GetHead() *Node {
	return list.Head
}
