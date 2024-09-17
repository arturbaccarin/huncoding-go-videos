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

func (q *Queue) Dequeue() string {
	if q.IsEmpty() {
		return ""
	}

	temp := q.front
	q.front = q.front.next
	temp.next = nil
	q.length--
	return temp.data
}

func (q *Queue) generateBinaryNumbers(n int) []string {
	result := make([]string, n)

	q.Enqueue("1")

	for i := 0; i < n; i++ {
		result[i] = q.Dequeue()
		n1 := result[i] + "0"
		n2 := result[i] + "1"
		q.Enqueue(n1)
		q.Enqueue(n2)
	}

	return result
}
