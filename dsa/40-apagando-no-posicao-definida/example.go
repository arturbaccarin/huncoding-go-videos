package main

import "fmt"

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) removeNodeAtPosition(position int) {
	if position < 0 {
		return
	}

	if position == 0 {
		ll.Head = ll.Head.Next
		return
	}

	count := 0
	previous := ll.Head
	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	previous.Next = current.Next
	current = nil
}

type Node struct {
	Data string
	Next *Node
}

func (n *Node) printList() {
	current := n
	for current != nil {
		fmt.Print(current.Data, "->")
		current = current.Next
	}
	fmt.Println("null")
}
