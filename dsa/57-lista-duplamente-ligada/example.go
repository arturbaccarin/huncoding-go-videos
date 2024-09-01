package main

import "fmt"

type DoublyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Data     string
	Previous *ListNode
	Next     *ListNode
}

func (dl *DoublyLinkedList) printList() {
	current := dl.Head

	for current != nil {
		fmt.Print("%d -> ", current.Data)
		current = current.Next
	}

	fmt.Print("null")
}
