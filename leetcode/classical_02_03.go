package leetcode

func deleteNode_(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
