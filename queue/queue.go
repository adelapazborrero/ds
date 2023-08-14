package queue

type Queue[T any] struct {
	Items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Add(itm T) {
	q.Items = append(q.Items, itm)
}

func (q *Queue[T]) Remove() {
	if length := len(q.Items); length > 0 {
		q.Items = q.Items[1:]
	}
}
