package leetcode

import (
	"testing"
)

func TestDLList(t *testing.T) {
	cache := NewLRUCache(2)
	list := cache.dllist

	if list.head.next != list.tail || list.tail.prev != list.head {
		t.Error("wrong constructor method")
	}

	node1 := &DLLNode{key: 1, val: 1}
	node2 := &DLLNode{key: 2, val: 2}
	node3 := &DLLNode{key: 3, val: 3}

	if !list.isEmpty() {
		t.Error("wrong isEmpty method")
	}

	list.addToHead(node1) // head <-> node1 <-> tail
	if list.head.next != node1 || node1.prev != list.head || node1.next != list.tail {
		t.Error("wrong addToHead method")
	}

	copyNode := list.last() // node1
	if copyNode.next != list.tail || list.tail.prev != copyNode {
		t.Error("wrong last method")
	}

	list.removeLast() // head <-> tail
	if list.head.next != list.tail || list.tail.prev != list.head {
		t.Error("wrong removeLast method")
	}

	list.addToHead(node2) // head <-> node2 <-> tail
	if list.head.next != node2 || node2.prev != list.head || node2.next != list.tail {
		t.Error("wrong addToHead method")
	}

	list.addToHead(node3) // head <-> node3 <-> node2 <-> tail
	if list.head.next != node3 || node3.prev != list.head || node3.next != node2 || node2.prev != node3 || node2.next != list.tail {
		t.Error("wrong addToHead method")
	}

	list.moveToHead(node2) // head <-> node2 <-> node3 <-> tail
	if list.head.next != node2 || node2.prev != list.head || node2.next != node3 || node3.prev != node2 || node3.next != list.tail {
		t.Error("wrong moveLastToHead method")
	}

	if list.size() != 2 {
		t.Error("wrong size method")
	}
}

func TestNewLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	if cache.capacity != 2 {
		t.Error("wrong constructor method")
	}

	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		t.Error("cache Get error")
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Error("cache Get error")
	}
	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Error("cache Get error")
	}
	if cache.Get(3) != 3 {
		t.Error("cache Get error")
	}
	if cache.Get(4) != 4 {
		t.Error("cache Get error")
	}

}

func TestNewLRUCache2(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put(2, 1)
	cache.Put(3, 2)

	if v := cache.Get(3); v != 2 {
		t.Errorf("error. got: %d", v)
	}

	if v := cache.Get(2); v != 1 {
		t.Errorf("error. got: %d", v)
	}
	cache.Put(4, 3)

	if v := cache.Get(2); v != 1 {
		t.Errorf("error. got: %d", v)
	}
	if v := cache.Get(3); v != -1 {
		t.Errorf("error. got: %d", v)
	}
	if v := cache.Get(4); v != 3 {
		t.Errorf("error. got: %d", v)
	}

}
