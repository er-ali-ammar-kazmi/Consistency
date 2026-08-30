package sorting

import (
	"cmp"
)

// O(n2)
func InsertionSort[T cmp.Ordered](arr *[]T) {

	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		count := 0
		j := i - 1
		for ; j >= 0 && (*arr)[j] > key; j-- {
			count++
			(*arr)[j+1] = (*arr)[j]
		}
		(*arr)[j+1] = key
	}
}
