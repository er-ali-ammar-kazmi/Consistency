package structure

type Queue[T any] []T

func (q *Queue[T]) Push(ele T) bool {
	if q.IsFull() {
		return false
	}
	*q = append(*q, ele)
	return true
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	ele := (*q)[0]
	*q = (*q)[1:]
	return ele, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) IsFull() bool {
	return len(*q) == cap(*q)
}
