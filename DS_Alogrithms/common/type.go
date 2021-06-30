package common

type ElementType int

type Node struct {
	val  ElementType
	next *Node
}

type TwoPrtNode struct {
	val  ElementType
	next *TwoPrtNode
	prev *TwoPrtNode
}

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}
