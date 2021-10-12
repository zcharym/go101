package queue

import (
	"github.com/zchary-ma/go101/DS_Alogrithms/common"
)

type Queue interface {
	EnQueue(value interface{})
	DeQueue() (interface{}, bool)
	Front() interface{}
	Full() bool
	common.Container
}
