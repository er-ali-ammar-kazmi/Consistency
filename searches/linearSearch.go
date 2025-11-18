package search

import "fmt"

func LinearSearch(arr []int, target int) (int, bool) {
	fmt.Println(arr, target)
	for idx, ele := range arr {
		if ele == target {
			return idx, true
		} else if ele > target {
			return idx, false
		}
	}
	return -1, false
}
