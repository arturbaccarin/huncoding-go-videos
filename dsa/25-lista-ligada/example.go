package main

import "fmt"

/*
guarda nós ligados
sequência de nós
nós guarda informação valor e endereço do próximo nó
primeiro nó = head
último nó aponta nulo
*/

type Node struct {
	data string
	next *Node
}

func (n *Node) printList() {
	current := n
	for current != nil {
		fmt.Print(current.data, "->")
		current = current.next
	}
	fmt.Println("null")
}

func (n *Node) findLength() int {
	count := 0
	current := n

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func main() {
	nodes := Node{
		data: "1",
		next: &Node{
			data: "2",
			next: &Node{
				data: "3",
				next: nil,
			},
		},
	}

	nodes.printList()
	fmt.Println(nodes.findLength())
}

// errado
func (n *Node) insertAtBeginning(data string) {
	newNode := &Node{data: data, next: n}
	n = newNode
	// Assignment to the method receiver propagates only to calles but no to callers
}
