package queue

import (
	"go101/DS_Alogrithms/common"
)

type Queue interface {
	EnQueue(value interface{})
	DeQueue() (interface{}, bool)
	Front() interface{}
	Full() bool
	common.Container
}
