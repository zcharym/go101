package stack

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := New()
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// if actualValue := stack.Values(); actualValue[0].(int) != 3 || actualValue[1].(int) != 2 || actualValue[2].(int) != 1 {
	// 	t.Errorf("Got %v expected %v", actualValue, "[3,2,1]")
	// }
	if actualValue := stack.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := stack.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := stack.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestStackPeek(t *testing.T) {
	stack := New()
	if actualValue, ok := stack.Peek(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if actualValue, ok := stack.Peek(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestStackPop(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	if actualValue, ok := stack.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := stack.Pop(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := stack.Pop(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue, ok := stack.Pop(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := stack.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	// if actualValue := stack.Values(); len(actualValue) != 0 {
	// 	t.Errorf("Got %v expected %v", actualValue, "[]")
	// }
}
