package main

/*
string - "{()}" - true
string - "{]" - false
string - "{()" - false
*/

type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) peek() int {
	if len(s.items) == 0 {
		return 0
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) pop() int {
	if len(s.items) == 0 {
		return 0
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func nextGreaterElement(arr []int) []int {
	n := len(arr)
	res := make([]int, n)

	stack := Stack{make([]int, 0)}

	for i := n - 1; i >= 0; i-- {
		for !stack.isEmpty() {
			for !stack.isEmpty() && stack.peek() <= arr[i] {
				stack.pop()
			}
		}

		if stack.isEmpty() {
			res[i] = -1
		} else {
			res[i] = stack.peek()
		}

		stack.push(arr[i])
	}

	return res
}
