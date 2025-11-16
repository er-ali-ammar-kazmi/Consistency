package sorting

import "fmt"

func BubbleSort(arr []int) {
	for i, _ := range arr {
		swapped := false
		count := 0
		for idx := 0; idx < len(arr)-i-1; idx++ {
			if arr[idx] > arr[idx+1] {
				arr[idx], arr[idx+1] = arr[idx+1], arr[idx]
				swapped = true
			}
			count++
		}
		fmt.Println(arr, count)
		if !swapped {
			break
		}
	}
}
