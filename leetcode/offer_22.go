package leetcode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	front, rear := head, head
	for ; k > 0; k-- {
		rear = rear.Next
	}
	for rear != nil {
		front, rear = front.Next, rear.Next
	}
	return front
}
