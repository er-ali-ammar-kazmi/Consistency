package structure

type Stack[T any] []T

func (s *Stack[T]) Push(ele T) bool {
	if s.IsFull() {
		return false
	}
	*s = append(*s, ele)
	return true
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	idx := len(*s) - 1
	ele := (*s)[idx]
	*s = (*s)[:idx]
	return ele, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) IsFull() bool {
	return len(*s) == cap(*s)
}
