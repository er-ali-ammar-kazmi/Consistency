package program

import "practise/structure"

func StackReverse[T any](stack *structure.Stack[T]) {
	if stack.IsEmpty() {
		return
	}
	top, _ := stack.Pop()
	StackReverse(stack)
	AddBefore(stack, top)
}

func AddBefore[T any](stack *structure.Stack[T], ele T) {
	if stack.IsEmpty() {
		stack.Push(ele)
		return
	}
	top, _ := stack.Pop()
	AddBefore(stack, ele)
	stack.Push(top)
}
