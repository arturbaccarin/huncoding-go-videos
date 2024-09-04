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

func (dl *DoublyLinkedList) removeNodeAtBeginning() {
	if dl.Lenght == 0 {
		return
	}

	temp := dl.Head
	if dl.Head == dl.Tail {
		dl.Tail = nil
	} else {
		dl.Head.Next.Previous = nil
	}

	dl.Head = dl.Head.Next
	temp.Next = nil
}
