package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  10,
				Next: nil,
			},
		},
	}

	ll := mergeTwoLists(l1, l2)

	current := ll
	for current != nil {
		fmt.Print(current.Val, "->")
		current = current.Next
	}

	fmt.Printf("null\n")
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	mergedList := &ListNode{}
	tail := mergedList

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			tail.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}

		tail = tail.Next
	}

	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}

	return mergedList.Next
}
