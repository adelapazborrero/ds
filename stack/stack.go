package stack

type Stack[T any] struct {
	Items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Add(itm T) {
	s.Items = append(s.Items, itm)
}

func (s *Stack[T]) Remove() {
	if length := len(s.Items); length > 0 {
		s.Items = s.Items[:length-1]
	}
}
