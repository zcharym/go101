package leetcode

/**
LRU (Least Recently Used) 算法就是一种缓存淘汰策略
计算机的缓存容量有限，如果缓存满了就要删除一些内容，给新内容腾位置。

当前实现：hash map + 双向链表
1. 定义一个带有头尾节点的双向链表，每个节点存储一个key-value对，链表的头尾指向最近使用的节点
2. 当缓存不满，就添加key-value对到链表头部
3. 当缓存满了，就删除链表的尾部节点，并将key-value对添加到链表头部
4. 当访问某个key时，将该节点移动到链表头部
5. 当缓存中有key时，直接返回value, 否则返回-1
*/

type LRUCache struct {
	capacity int
	cache    map[int]*DLLNode
	dllist   *DLList
}

type DLList struct {
	head  *DLLNode
	tail  *DLLNode
	count int
}

type DLLNode struct {
	key  int
	val  int
	prev *DLLNode
	next *DLLNode
}

func NewLRUCache(capacity int) LRUCache {
	dummyHead := &DLLNode{key: -1, val: -1}
	dummyTail := &DLLNode{key: -1, val: -1}
	dummyTail.prev = dummyHead
	dummyHead.next = dummyTail

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLLNode, capacity),
		dllist: &DLList{
			head: dummyHead,
			tail: dummyTail,
		}}
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.cache[key]; ok {
		c.dllist.moveToHead(v)
		return v.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if v, ok := c.cache[key]; ok {
		c.dllist.moveToHead(v)
		v.val = value
	} else {
		if c.dllist.count == c.capacity {
			delete(c.cache, c.dllist.tail.prev.key)
			c.dllist.removeLast()
		}
		node := &DLLNode{key: key, val: value}
		c.cache[key] = node
		c.dllist.addToHead(node)
	}
}

func (l *DLList) addToHead(node *DLLNode) {
	node.next = l.head.next
	node.next.prev = node
	node.prev = l.head
	l.head.next = node
	l.count++
}

func (l *DLList) last() *DLLNode {
	return l.tail.prev
}

func (l *DLList) removeLast() {
	if l.isEmpty() {
		return
	}
	l.tail.prev.prev.next = l.tail
	l.tail.prev = l.tail.prev.prev
	l.count--
}

func (l *DLList) moveToHead(node *DLLNode) {
	if l.size() < 2 {
		return
	}

	if l.head.next == node {
		return
	}
	l.removeNode(node)
	l.addToHead(node)
}

func (l *DLList) removeNode(node *DLLNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.count--
}

func (l *DLList) isEmpty() bool {
	return l.head.next == l.tail
}

func (l *DLList) size() int {
	return l.count
}
