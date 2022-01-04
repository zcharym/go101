package stack

import (
	"github.com/zchary-ma/impl-ds/common"
)

type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)
	common.Container
}
