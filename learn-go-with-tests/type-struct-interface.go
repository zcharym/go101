package learning_tests

type Rectangle struct {
	width, height float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}

func Area(rect Rectangle) float64 {
	return rect.width * rect.height
}
