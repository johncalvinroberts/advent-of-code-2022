package utils

import "fmt"

type QueueNode[T comparable] struct {
	value T
	next  *QueueNode[T]
}

type Queue[T comparable] struct {
	length int
	head   *QueueNode[T]
	tail   *QueueNode[T]
}

func (q *Queue[T]) Enqueue(value T) {
	n := &QueueNode[T]{value: value}
	if q.length == 0 {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
	q.length++
}

func (q *Queue[T]) Dequeue() *QueueNode[T] {
	var ret *QueueNode[T]
	if q.head != nil {
		ret = q.head
		q.head = ret.next
		q.length--
	} else {
		fmt.Println("The queue is empty.")
	}
	return ret
}

func (q *Queue[T]) Peek() T {
	return q.head.value
}
