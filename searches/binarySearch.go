package search

import (
	"cmp"
	"fmt"
)

func BinarySearch[T cmp.Ordered](arr []T, target T) (int, bool) {
	var left, mid, right = 0, 0, len(arr) - 1

	for left <= right {
		mid = (left + right) / 2
		fmt.Println(mid)
		if arr[mid] == target {
			return mid, true
		}

		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1, false
}
