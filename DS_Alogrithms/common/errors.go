package common

type RangeError struct {
}

type InputError struct {
}

func (i InputError) Error() string {
	return ">>>input error"
}

func (r RangeError) Error() string {
	return ">>>index out of range"
}
