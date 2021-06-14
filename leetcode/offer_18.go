package leetcode

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	pre := head
	for pre.Next != nil && pre.Next.Val != val {
		pre = pre.Next
	}

	if pre.Next != nil {
		pre.Next = pre.Next.Next
	}
	return head
}
