package leetcode

func detectCycle(head *ListNode) *ListNode {
	m := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}
