package queue

// CircularQueue 循环队列的顺序实现
type CircularQueue struct {
	data        []interface{}
	front, rear int
	size        int
}

func NewCircularQueue(size int) *CircularQueue {
	// 多加一个单元，用于rear与front判断
	q := make([]interface{}, size)
	return &CircularQueue{
		data:  q,
		front: 0,
		rear:  0,
		size:  0,
	}
}

func (q *CircularQueue) EnQueue(value interface{}) {
	if q.Full() {
		return
	}
	q.data[q.rear] = value
	q.rear = (q.rear + 1) % len(q.data)
	q.size++
}

func (q *CircularQueue) DeQueue() (interface{}, bool) {
	if q.Empty() {
		return nil, false
	}
	value := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	q.size--
	return value, true
}

func (q *CircularQueue) Front() interface{} {
	if q.Empty() {
		return nil
	}
	return q.data[q.front]
}

func (q *CircularQueue) Full() bool {
	if q.size == len(q.data) {
		return true
	}
	return false
}

func (q *CircularQueue) Empty() bool {
	if q.rear == q.front && q.size == 0 {
		return true
	} else {
		return false
	}
}

func (q *CircularQueue) Size() int {
	return q.size
}

func (q *CircularQueue) Clear() {
	q.rear = 0
	q.front = 0
	q.size = 0
	q.data = []interface{}{}
}

func (q *CircularQueue) Values() []interface{} {
	return q.data[:len(q.data)]
}
