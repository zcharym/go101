package stack

import (
	"go101/DS_Alogrithms/common"
)

type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)
	common.Container
}
