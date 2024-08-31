package main

import "fmt"

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
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	ll := addTwoNumbers(l1, l2)

	current := ll
	for current != nil {
		fmt.Print(current.Val, "->")
		current = current.Next
	}

	fmt.Printf("null\n")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tails := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		x := func() int {
			if l1 != nil {
				return l1.Val
			}

			return 0
		}()

		y := func() int {
			if l2 != nil {
				return l2.Val
			}

			return 0
		}()

		sum := carry + x + y
		carry = sum / 10
		tails.Next = &ListNode{Val: sum % 10}
		tails = tails.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		tails.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
