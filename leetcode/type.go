package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListQueue struct {
	front, end *ListNode
}
