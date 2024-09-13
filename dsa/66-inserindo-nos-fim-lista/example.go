package main

type CircularLinkedList struct {
	Length int
	Last   *ListNode
}

type ListNode struct {
	Data int
	Next *ListNode
}

func (cl *CircularLinkedList) InsertNodeAtEnd(data int) {
	newNode := &ListNode{Data: data}

	if cl.Length == 0 || cl.Last == nil {
		cl.Last = newNode
		cl.Last.Next = cl.Last
	} else {
		newNode.Next, cl.Last.Next = cl.Last.Next, newNode
		cl.Last = newNode
	}

	cl.Length++
}
