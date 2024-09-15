package main

import "fmt"

/*
string - "{()}" - true
string - "{]" - false
string - "{()" - false
*/

type Stack struct {
	items []string
}

func (s *Stack) push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) peek() string {
	if len(s.items) == 0 {
		return ""
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) pop() string {
	if len(s.items) == 0 {
		return ""
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func isValid(s string) bool {
	stack := Stack{}

	for _, charPos := range s {
		char := fmt.Sprintf("%c", charPos)

		if char == "(" || char == "{" || char == "[" {
			stack.push(char)
		} else {
			if stack.isEmpty() {
				return false
			}

			top := stack.peek()
			if (char == ")" && top == "(") || (char == "}" && top == "{") || (char == "]" && top == "[") {
				stack.pop()
			} else {
				return false
			}
		}
	}

	return stack.isEmpty()
}
