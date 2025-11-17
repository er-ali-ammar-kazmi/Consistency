package sorting

import "fmt"

func BubbleSort(arr []int) {
	fmt.Println(arr)
	for i, _ := range arr {
		swapped := false
		count := 0
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
			count++
		}
		fmt.Println(arr, count)
		if !swapped {
			break
		}
	}
	fmt.Println(arr)
}
