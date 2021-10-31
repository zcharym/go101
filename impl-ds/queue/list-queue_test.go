package queue

import (
	"reflect"
	"testing"
)

func TestCircularQueue_EnQueue(t *testing.T) {
	queue := NewCircularQueue(4)
	queue.EnQueue(2)
	queue.EnQueue("3")

	if queue.rear != 2 {
		t.Errorf("rear pointer error:expect 2, got %d", queue.rear)
	}

	queue.EnQueue("3")
	queue.EnQueue("4")
	checked := make([]interface{}, 0)
	checked = append(checked, 2, "3", "3", nil)

	if !reflect.DeepEqual(queue.data, checked) {
		t.Errorf("error expect equal, checked: %v,got %v", checked, queue.data)
	}
}

func TestListQueue_All(t *testing.T) {
	queue := NewCircularQueue(4)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)

	checked := make([]interface{}, 0)
	checked = append(checked, 1, 2, 3, 4)

	if !reflect.DeepEqual(checked, queue.data) {
		errorMessage(t, queue)
	}

	if queue.Front() != 1 {
		errorMessage(t, queue)
	}

	if !queue.Full() {
		errorMessage(t, queue)
	}

	queue.DeQueue()
	queue.EnQueue(4)

	checked2 := make([]interface{}, 0)
	checked2 = append(checked2, 4, 2, 3, 4)
	if !reflect.DeepEqual(checked2, queue.data) {
		errorMessage(t, queue)
	}
}

func errorMessage(t *testing.T, q *CircularQueue) {
	t.Errorf("Error occured, front:%d, rear:%d, data:%v", q.front, q.rear, q.data)
}
