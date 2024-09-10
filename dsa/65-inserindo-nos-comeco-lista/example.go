package main

/*
Lista Ligada -> Um nó aponta para o próximo, mas não sabe quem é o nó que vem antes dele
Lista Duplamente Ligada -> um nó aponta para o próximo e o próximo aponta para o anterior
Lista Ligada Circular -> é uma lista ligada que o último nó aponta para o primeiro.

Se só tiver um item dentro dela, ele vai apontar para ele mesmo
*/

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
	cl := &CircularLinkedList{
		Length: 1,
		Last:   listNode,
	}
}
