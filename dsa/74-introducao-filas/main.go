package main

// FIFO
// duas variáveis, front/top e rear

// Enqueue: coloca o dado

type ListNode struct {
	data string
	next *ListNode
}

type Queue struct {
	length int
	front  *ListNode
	rear   *ListNode
}

func (ln *Queue) Length() int {
	return ln.length
}

func (ln *Queue) IsEmpty() bool {
	return ln.length == 0
}

func (q *Queue) Enqueue(data string) { // or offer and poll to remove
	temp := &ListNode{
		data: data}

	if q.IsEmpty() {
		q.front = temp
	} else {
		q.rear.next = temp
	}

	q.rear = temp
	q.length++
}
