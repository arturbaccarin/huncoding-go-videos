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
}
