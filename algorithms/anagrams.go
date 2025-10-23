package algorithms

import "maps"

func Anagrams(string1, string2 string) bool {
	// silent == listen
	if len(string1) != len(string2) {
		return false
	}

	string1Count := map[rune]int{}
	string2Count := map[rune]int{}
	for _, ch := range string1 {
		string1Count[ch]++
	}
	for _, ch := range string2 {
		string2Count[ch]++
	}
	return maps.Equal(string1Count, string2Count)
}
