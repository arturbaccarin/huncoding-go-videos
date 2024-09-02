package main

import "fmt"

type DoublyLinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Lenght int
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

func (dl *DoublyLinkedList) insertAtBeginning(data string) {
	newNode := &ListNode{
		Data: data,
	}

	if dl.Lenght == 0 {
		dl.Head = newNode
	} else {
		dl.Head.Previous = newNode
	}

	newNode.Next = dl.Head
	dl.Head = newNode
	dl.Lenght++
}
