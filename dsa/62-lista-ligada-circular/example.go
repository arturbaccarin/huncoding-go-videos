package main

type CircularLinkedList struct {
	Length int
	Last   *ListNode
}

type ListNode struct {
	Data int
	Next *ListNode
}

func main() {
	listNode := &ListNode{
		Data: 2,
		Next: listNode2,
	}

	listNode3 := &ListNode{
		Data: 3,
	}

	listNode2 := &ListNode{
		Data: 3,
		Next: listNode3,
	}

	listNode3.Next = listNode

	cl := &CircularLinkedList{
		Length: 1,
		Last:   listNode3,
	}
}
