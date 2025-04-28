// https://youtu.be/8JOrxl0KkDs
// 6:17
package main

import (
	"sync/atomic"
	"unsafe"
)

type Node struct {
	value int
	next  *Node
}

type LockFreeQueue struct {
	head *Node
	tail *Node
}

func NewLockFreeQueue() *LockFreeQueue {
	node := &Node{}
	return &LockFreeQueue{
		head: node,
		tail: node,
	}
}

func (q *LockFreeQueue) Enqueue(value int) {
	node := &Node{value: value}

	for {
		tail := q.tail
		next := tail.next

		if atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(&tail.next)),
			unsafe.Pointer(next),
			unsafe.Pointer(node),
		) {
			atomic.CompareAndSwapPointer(
				(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
				unsafe.Pointer(tail),
				unsafe.Pointer(node),
			)

			return
		}
	}
}

func (q *LockFreeQueue) Dequeue() (int, bool) {
	for {
		head := q.head
		tail := q.tail
		next := head.next

		if head == tail {
			if next == nil {
				return 0, false
			}

			atomic.CompareAndSwapPointer(
				(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
				unsafe.Pointer(tail),
				unsafe.Pointer(next),
			)
		} else {
			value := next.value
			if atomic.CompareAndSwapPointer(
				(*unsafe.Pointer)(unsafe.Pointer(&q.head)),
				unsafe.Pointer(head),
				unsafe.Pointer(next),
			) {
				return value, true
			}
		}
	}
}

func main() {
	queue := NewLockFreeQueue()

	go queue.Enqueue(1)
	go queue.Enqueue(2)
	go queue.Enqueue(3)

	value, ok := queue.Dequeue()
	if ok {
		println(value)
	} else {
		println("empty")
	}
}

/*
func main() {
	var wg sync.WaitGroup

	var counter int64

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1) // a demora é a mesma, mas com menos uso de máquina
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
*/

/*
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
*/
