package stack

type Stack interface {
	InitStack()
	IsEmpty() (ok bool)
	Push(value interface{})
	Pop(elem interface{})
	Peek(elem interface{})
}
