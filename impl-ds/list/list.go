package list

import "github.com/zchary-ma/impl-ds/common"

type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Insert(index int, values ...interface{})
	Set(index int, value interface{})
	IndexOf(value interface{}) int
	// Contains(values ...interface{}) bool
	// Swap(index1, index2 int)

	common.Container
}
