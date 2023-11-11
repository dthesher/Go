package dqueue

import "fmt"

type Queue struct {
	data  []int
	front int //front element ka index
	rear  int //rear element ka index
}

func NewQueue() *Queue {
	q := &Queue{
		data:  make([]int, 0),
		front: 0,
		rear:  -1,
	}
	return q
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
	q.rear++
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("empty result queue")
		return -1
	}

	value := q.data[q.front]
	for i := q.front; i < q.rear; i++ {
		q.data[i] = q.data[i+1]
	}

	q.data = q.data[0:q.rear]
	q.rear--
	return value
}

func (q *Queue) IsEmpty() bool {
	return q.front > q.rear
}

func (q *Queue) Show() {
	for _, i := range q.data {
		fmt.Printf("%v, ", i)
	}
	println()
}

func (q *Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.front]
}

func (q *Queue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.rear]
}
