package main

import "fmt"

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) checkLoop() bool {
	fastPointer := ll.Head
	slowPointer := ll.Head

	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if slowPointer == fastPointer {
			return true
		}
	}

	return false
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