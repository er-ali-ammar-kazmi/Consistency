package program

import (
	"fmt"
	"sort"
)

func SumsOfLeastTwo(w []int, n int) int {
	if len(w) < 2 {
		return 0
	}
	sort.Ints(w)
	S := w[0] + w[1]
	w = append(w[2:], S)
	S += SumsOfLeastTwo(w, n)
	return S
}

func SumOfTwos(nums []int, target int) {
	dict := map[int]int{}
	for idx, num := range nums {
		if id, ok := dict[target-num]; ok {
			fmt.Println([]int{id, idx})
		} else {
			dict[num] = idx
		}
	}
}
