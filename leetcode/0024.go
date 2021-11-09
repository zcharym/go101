package leetcode

func swapPairs(head *ListNode) *ListNode {
	cur := head
	if head != nil {
		for cur.Next != nil {
			cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
			if cur.Next.Next != nil {
				cur = cur.Next.Next
			} else {
				break
			}
		}
	}
	return head
}
