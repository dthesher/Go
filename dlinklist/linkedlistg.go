package dlinklist

import "fmt"

// Node represents a node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
}

// NewLinkedList initializes and returns a new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds a new node with the given data to the end of the linked list
func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Display prints the elements of the linked list
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// Search searches for a given value in the linked list
func (ll *LinkedList) Search(value int) *Node {
	current := ll.head
	for current != nil {
		if current.data == value {
			return current
		}
		current = current.next
	}
	return nil
}

// Delete removes the first occurrence of a value from the linked list
func (ll *LinkedList) Delete(value int) {
	if ll.head == nil {
		return
	}

	if ll.head.data == value {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next == nil {
		return
	}

	current.next = current.next.next
}

// InsertAfter inserts a new node with the given data after a specific node
func (ll *LinkedList) InsertAfter(prevNode *Node, data int) {
	if prevNode == nil {
		fmt.Println("Previous node cannot be nil")
		return
	}

	newNode := &Node{data: data, next: prevNode.next}
	prevNode.next = newNode
}

// Reverse reverses the linked list in-place
func (ll *LinkedList) Reverse() {
	var prev *Node = nil
	current := ll.head
	var next *Node = nil

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	ll.head = prev
}

// Sort performs a simple bubble sort on the linked list
func (ll *LinkedList) Sort() {
	if ll.head == nil {
		return
	}

	swapped := true
	for swapped {
		swapped = false
		current := ll.head
		for current.next != nil {
			if current.data > current.next.data {
				current.data, current.next.data = current.next.data, current.data
				swapped = true
			}
			current = current.next
		}
	}
}

// GetHead returns the head of the linked list
func (list *LinkedList) GetHead() *Node {
	return list.head
}
