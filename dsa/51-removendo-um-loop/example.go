package main

import "fmt"

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) findWhereLoopStarts() *Node {
	fastPointer := ll.Head
	slowPointer := ll.Head

	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if slowPointer == fastPointer {
			temp := ll.Head
			for slowPointer != temp {
				temp = temp.Next
				slowPointer = slowPointer.Next
			}

			slowPointer.Next = nil
		}
	}

	return nil
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
