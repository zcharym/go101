package queue

// ListQueue 循环队列的顺序实现
type ListQueue struct {
	data        []interface{}
	front, rear int
}

func New(size int) *ListQueue {
	q := make([]interface{}, size)
	return &ListQueue{
		data:  q,
		front: 0,
		rear:  0,
	}
}

func (q *ListQueue) EnQueue(value interface{}) {
	size := len(q.data)
	if (q.rear+1)%size == q.front {
		return
	}
	q.data[q.rear] = value
	q.rear = (q.rear + 1) % size
}

func (q *ListQueue) DeQueue() (interface{}, bool) {
	size := len(q.data)
	if q.rear == q.front {
		return nil, false
	}
	value := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % size
	return value, true
}

func (q *ListQueue) Empty() bool {
	if q.rear == q.front {
		return true
	} else {
		return false
	}
}

func (q *ListQueue) Size() int {
	return len(q.data)
}

func (q *ListQueue) Clear() {
	q.rear = 0
	q.front = 0
	q.data = []interface{}{}
}

func (q *ListQueue) Values() []interface{} {
	return q.data
}
