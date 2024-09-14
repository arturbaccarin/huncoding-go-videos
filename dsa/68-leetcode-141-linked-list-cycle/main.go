package main

// definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fastPointer := head
	slowPointer := head

	if fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if fastPointer == slowPointer {
			return true
		}
	}

	return false
}
