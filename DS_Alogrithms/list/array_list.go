package list

type SqList struct {
	data   []interface{}
	length int
}

func New(values ...interface{}) *SqList {
	sqList := &SqList{}
	if len(values) > 0 {
		sqList.Add(values...)
	}
	return sqList
}

func (l *SqList) IndexOf(value interface{}) int {
	if l.length > 0 {
		for index, elem := range l.data {
			if elem == value {
				return index
			}
		}
	}
	return -1
}

func (l *SqList) Add(values ...interface{}) {
	l.growBy(len(values))
	for _, value := range values {
		l.data[l.length] = value
		l.length++
	}
}

func (l *SqList) Get(index int) (interface{}, bool) {
	if l.inRange(index) {
		return l.data[index], true
	}
	return nil, false
}

func (l *SqList) Remove(index int) {
	if !l.inRange(index) {
		return
	}
	// AMEND c-style remove, not appropriate in go
	for i := index; i < l.length-2; i++ {
		l.data[i] = l.data[i+1]
	}
	l.data = l.data[:len(l.data)-1]
	l.length--
}

func (l *SqList) Insert(index int, values ...interface{}) {
	if !l.inRange(index) {
		if index == l.length {
			l.Add(values...)
		}
		return
	}
	length := len(values)
	l.growBy(length)
	l.length += length
	copy(l.data[index+length:], l.data[index:l.length-length])
	copy(l.data[index:], values)
}

func (l *SqList) Set(index int, value interface{}) {
	if !l.inRange(index) {
		if index == l.length {
			l.Add(value)
		}
		return
	}
	l.data[index] = value
}

func (l *SqList) Empty() bool {
	return l.length == 0
}

func (l *SqList) Size() int {
	return l.length
}

func (l *SqList) Clear() {
	l.length = 0
	l.data = []interface{}{}
}

func (l *SqList) Values() []interface{} {
	newList := make([]interface{}, l.length, l.length)
	// NOTE
	copy(newList, l.data[:l.length])
	return newList
}

func (l *SqList) inRange(index int) bool {
	return index >= 0 && index < l.length
}

func (l *SqList) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(l.data)
	if l.length+n >= currentCapacity {
		newCapacity := int(float32(2.0) * float32(currentCapacity+n))
		newElements := make([]interface{}, newCapacity, newCapacity)
		copy(newElements, l.data)
		l.data = newElements
	}
}
