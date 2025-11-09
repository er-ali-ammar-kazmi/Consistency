package program

import (
	"math"
)

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func NextPrime(num int) int {
	if num < 2 {
		return 2
	}
	for i := num + 1; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
}
