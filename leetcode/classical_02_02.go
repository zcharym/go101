package leetcode

// kthToLastV1 快慢指针
func kthToLastV1(head *ListNode, k int) int {
	slow := head
	fast := head
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

// kthToLastV2 通过迭代反转链表，
func kthToLastV2(head *ListNode, k int) int {
	// TODO
	return 0
}

// kthToLastV3 数组缓存（内存占用过多）
func kthToLastV3(head *ListNode, k int) int {
	var (
		s      []int
		length int
	)
	for head != nil {
		s = append(s, head.Val)
		length++
		head = head.Next
	}
	return s[length-k]
}
