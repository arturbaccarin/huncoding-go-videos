package main

import "fmt"

/*
guarda nós ligados
sequência de nós
nós guarda informação valor e endereço do próximo nó
primeiro nó = head
último nó aponta nulo
*/

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) InsertFirst(data string) {
	newNode := &Node{Data: data, Next: ll.Head} // o cara da primeira posição vira o da segunda posição
	ll.Head = newNode
}

type Node struct {
	Data string
	Next *Node
}

func (n *Node) printList() {
	current := n
	for current != nil {
		fmt.Print(current.Data, "->")
		current = current.Next
	}
	fmt.Println("null")
}

func main() {
	nodes := LinkedList{}
	nodes.InsertFirst("world")
	nodes.InsertFirst("hello")

	nodes.Head.printList()
}
