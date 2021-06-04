package learning_tests

import (
	"errors"
)

type Dictionary map[string]string

var ErrorNotFound = errors.New("can't find the word")

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return word, nil
}
