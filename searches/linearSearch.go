package search

func LinearSearch[T comparable](arr []T, target T) (int, bool) {
	for idx, ele := range arr {
		if ele == target {
			return idx, true
		}
	}
	return -1, false
}
