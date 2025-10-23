package structure

type Stack []int

func (s *Stack) Push(ele int) {
	*s = append(*s, ele)
}

func (s *Stack) Pop() int {
	idx := len(*s) - 1
	ele := (*s)[idx]
	*s = (*s)[:idx]
	return ele
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) IsFull() bool {
	return len(*s) == cap(*s)
}
