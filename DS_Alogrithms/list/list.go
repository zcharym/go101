package list

type List interface {
	Get(index int) (elementType, error)
	Remove(index int) error
	Add(values ...elementType) error
	Insert(index int, value elementType) error
	FindByValue(elementType) (index int, err error)
	Length() (length int)
}
