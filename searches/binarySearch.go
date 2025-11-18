package search

import "fmt"

func BinarySearch(arr []int, target int) (int, bool) {
	fmt.Println(arr, target)
	if target > arr[len(arr)-1] {
		return -1, false
	}

	var left, mid, right = 0, 0, len(arr) - 1

	for left <= right {
		mid = (left + right) / 2
		fmt.Println(left, mid, right)

		if arr[mid] == target {
			return mid, true
		}

		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return mid, false
}
