package program

import (
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
