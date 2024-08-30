package main

import "fmt"

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) insertAtGivenPosition(data string, position int) {
	newNode := &Node{Data: data}

	if position == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	previous := ll.Head
	count := 0

	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	newNode.Next = current
	previous.Next = newNode

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
