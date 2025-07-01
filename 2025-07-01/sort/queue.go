package sort

type QueueNode struct {
	Value int
	Next  *QueueNode
}

type Queue struct {
	front *QueueNode
	rear  *QueueNode
	size  int
}

// Enqueue adds a value to the back of the queue.
func (q *Queue) Enqueue(value int) {
	newNode := &QueueNode{Value: value}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}

	q.size++
}

// Dequeue removes and returns the front value.
func (q *Queue) Dequeue() (int, bool) {
	if q.front == nil {
		return 0, false
	}

	val := q.front.Value
	q.front = q.front.Next
	q.size--

	if q.front == nil {
		q.rear = nil // queue is empty
	}

	return val, true
}

// Peek returns the front value without removing it.
func (q *Queue) Peek() (int, bool) {
	if q.front == nil {
		return 0, false
	}
	return q.front.Value, true
}

// IsEmpty checks if the queue has no elements.
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Len returns the number of elements in the queue.
func (q *Queue) Len() int {
	return q.size
}
