package program

import "iter"

func Fibonacci(num int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for a <= num {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}
