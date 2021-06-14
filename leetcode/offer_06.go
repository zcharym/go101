package leetcode

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	// NOTE for loop to reverse array
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}
