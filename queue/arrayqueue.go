package queue

import "errors"

type ArrayQueue struct {
	elements []interface{}
	size  int
	front int
	rear  int
}

func assertArrayQueueImplementation() {
	var _ Queue = (*ArrayQueue)(nil)
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		elements: make([]interface{}, capacity),
		size:     0,
		front:	  0,
		rear:	  0,
	}
}

func NewDefault() *ArrayQueue {
	return &ArrayQueue{
		elements: make([]interface{}, 10),
		size:     0,
		front:	  0,
		rear:	  0,
	}
}

func (q *ArrayQueue) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *ArrayQueue) GetSize() int {
	return q.size
}

func (q *ArrayQueue) Enqueue(e interface{}) error {
	if q.size == cap(q.elements) {
		return errors.New("queue is full")
	}
	q.elements[q.rear] = e
	q.size ++
	q.rear = (q.rear + 1) % cap(q.elements)
	return nil
}

func (q *ArrayQueue) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	q.size --
	e := q.elements[q.front]
	q.front = (q.front + 1) % cap(q.elements)
	return e, nil
}