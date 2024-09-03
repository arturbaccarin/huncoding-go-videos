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
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}

	fmt.Print("null")
}

func (dl *DoublyLinkedList) insertAtEnd(data string) {
	newNode := &ListNode{
		Data: data,
	}

	if dl.Lenght == 0 {
		dl.Head = newNode
	} else {
		dl.Tail.Next = newNode
		newNode.Previous = dl.Tail
	}

	dl.Tail = newNode
	dl.Lenght++
}
