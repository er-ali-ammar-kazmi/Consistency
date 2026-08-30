package sorting

import (
	"cmp"
)

// O(n2)
func BubbleSort[T cmp.Ordered](arr *[]T) {
	var swapped bool
	var count int
	for i, _ := range *arr {
		swapped = false
		count = 0
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swapped = true
			}
			count++
		}
		if !swapped {
			break
		}
	}
}
