package learning_tests

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", errors.New("can't find the word")
	}
	return word, nil
}
