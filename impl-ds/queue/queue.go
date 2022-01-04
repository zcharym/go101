package queue

import (
	"github.com/zchary-ma/impl-ds/common"
)

type Queue interface {
	EnQueue(value interface{})
	DeQueue() (interface{}, bool)
	Front() interface{}
	Full() bool
	common.Container
}
