package list

import (
	"go101/DS_Alogrithms/common"
)

type elementType int

type SqList struct {
	data   []elementType
	length int
}

func (s *SqList) Add(values ...elementType) error {
	s.length += len(values)
	s.data = append(s.data, values...)
	return nil
}

func (s *SqList) Insert(index int, value elementType) error {
	if index < 0 || index > s.length {
		return common.InputError{}
	}
	for i := s.length - 1; i > index; i-- {
		i := i
		s.data[i-1] = s.data[i]
	}
	s.data[index] = value
	s.length++
	return nil
}

func (s SqList) FindByValue(elem elementType) (index int, err error) {
	for i := 0; i < s.length; i++ {
		if s.data[i] == elem {
			return i, nil
		}
	}
	return -1, nil
}

// Get TODO Exported method with the unexported return type
func (s SqList) Get(index int) (elementType, error) {
	if index < 0 || index >= s.length {
		return 0, common.InputError{}
	}
	if s.length > 0 {
		return s.data[index], nil
	}
	return 0, common.RangeError{}
}

func (s *SqList) Remove(index int) error {
	if index < 0 || index >= s.length {
		return common.InputError{}
	}
	for i := index; i < s.length-1; i++ {
		s.data[i] = s.data[i+1]
	}
	// change to initial value
	s.data[s.length-1] = 0
	return nil
}

func (s SqList) Length() (length int) {
	return s.length
}
