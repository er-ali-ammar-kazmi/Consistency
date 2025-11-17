package sorting

import "fmt"

func SelectionSort(arr []int) {
	fmt.Println(arr)
	for i, _ := range arr {
		count := 0
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			count++
		}
		fmt.Println(arr, count)
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	fmt.Println(arr)
}
