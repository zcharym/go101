package leetcode

func deleteNodeBySelf(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
