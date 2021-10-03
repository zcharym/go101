package leetcode

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	m := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if !m[cur.Val] {
			m[cur.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}
