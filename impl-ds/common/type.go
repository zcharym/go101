package common

type ElementType int

type Node struct {
	Val  ElementType
	Next *Node
}

type TwoPrtNode struct {
	Val  ElementType
	Next *TwoPrtNode
	Prev *TwoPrtNode
}

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}
