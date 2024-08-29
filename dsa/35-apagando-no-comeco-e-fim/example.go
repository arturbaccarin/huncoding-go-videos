package main

import "fmt"

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) removeFirstNode() {
	if ll.Head == nil {
		return
	}

	ll.Head = ll.Head.Next
	return
}

func (ll *LinkedList) removeNodeAtEnd() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	previous := ll.Head
	for previous.Next.Next != nil {
		previous = previous.Next
	}

	previous.Next = nil
	return
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
