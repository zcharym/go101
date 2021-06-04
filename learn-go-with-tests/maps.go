package learning_tests

import (
	"errors"
)

type Dictionary map[string]string

var ErrorNotFound = errors.New("can't find the word")
var ErrorAddedAlready = errors.New("word already added")

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return word, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, ok := d[key]
	if ok {
		return ErrorAddedAlready
	}
	d[key] = value
	return nil
}
