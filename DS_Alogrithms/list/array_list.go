package list

type SqList struct {
	data   []interface{}
	length int
}

func New(values ...interface{}) *SqList {
	sqList := &SqList{}
	if len(values) > 0 {
		sqList.Add(values)
	}
	return sqList
}

func (l *SqList) Add(values ...interface{}) {
	l.data = append(l.data, values...)
	l.length += len(values)
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
	for i := l.length; i > index; i-- {
		l.data[i-1] = l.data[i]
	}
	l.data = l.data[:len(l.data)-1]
	l.length--
}

func (l *SqList) Insert(index int, values ...interface{}) {
	if !l.inRange(index) {
		if index == l.length {
			l.Add(values)
		}
		return
	}
	l.data = append(l.data, values...)
	l.length += len(values)
}

func (l *SqList) Set(index int, value interface{}) {
	if !l.inRange(index) {
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
	copy(newList, l.data[:l.length])
	return newList
}

func (l *SqList) inRange(index int) bool {
	return index >= 0 && index < l.length
}
