package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    p1,p2 := l1,l2
	for p1.Next != nil  && p2.Next != nil{
        if l2.Val <= l1.Val {
            node := &ListNode{}
            node.Val = l2.Val
            node.Next = l1
        }
		p1 = p1.Next
        p2 = p2.Next
	}

    if p1.Next == nil {
        p1.Next = p2
    }

    return l1
}
