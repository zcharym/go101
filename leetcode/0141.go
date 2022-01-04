package leetcode

// hasCycle set存址判重
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true
		} else {
			m[cur] = true
		}
		cur = cur.Next
	}
	return false
}

// hasCycleV2 快慢指针
func hasCycleV2(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
