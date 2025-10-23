package problems

import "dsa/structure"

func StackReverse(stack *structure.Stack) {
	if stack.IsEmpty() {
		return
	}
	top := stack.Pop()
	StackReverse(stack)
	AddBefore(stack, top)
}

func AddBefore(stack *structure.Stack, ele int) {
	if stack.IsEmpty() {
		stack.Push(ele)
		return
	}
	top := stack.Pop()
	AddBefore(stack, ele)
	stack.Push(top)
}
