package learning_tests

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrorNotFound     = DictionaryErr("can't find the word")
	ErrorAddedAlready = DictionaryErr("word already added")
)

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return word, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorNotFound:
		d[key] = value
	case nil:
		return ErrorAddedAlready
	default:
		return err
	}
	return nil
}
