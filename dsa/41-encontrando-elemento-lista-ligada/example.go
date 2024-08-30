package main

import "fmt"

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) findElement(element string) int {
	current := ll.Head
	count := 0

	for current != nil {
		if element == current.Data {
			return count
		}

		count++
		current = current.Next
	}

	return -1
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
