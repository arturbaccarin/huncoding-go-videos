package main

type CircularLinkedList struct {
	Length int
	Last   *ListNode
}

type ListNode struct {
	Data int
	Next *ListNode
}

func (cl *CircularLinkedList) InsertFirst(data int) {
	newNode := &ListNode{Data: data}

	if cl.Length == 0 || cl.Last == nil {
		cl.Last = newNode
	} else {
		newNode.Next = cl.Last.Next
	}

	cl.Last.Next = newNode
	cl.Length++
}

func main() {
	listNode := &ListNode{
		Data: 2,
		Next: nil,
	}

	// não usa o head
	_ = &CircularLinkedList{
		Length: 1,
		Last:   listNode,
	}
}
