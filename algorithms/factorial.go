package algorithms

func Fact(num int) int {
	// n! = n * n-1 * ...1
	if num <= 1 {
		return 1
	}
	return num * Fact(num-1)
}
