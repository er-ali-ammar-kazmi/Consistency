package sorting

import (
	"cmp"
)

// O(n2)
func SelectionSort[T cmp.Ordered](arr *[]T) {

	for i, _ := range *arr {
		minIndex := i
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			(*arr)[i], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[i]
		}
	}
}
