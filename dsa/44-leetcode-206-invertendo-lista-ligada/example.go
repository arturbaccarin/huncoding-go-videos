package main

import "fmt"

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) reverseList() {
	current := ll.Head

	var previous *Node
	var next *Node

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	ll.Head = previous
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
