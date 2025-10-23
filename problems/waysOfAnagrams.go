package problems

import algorithms "dsa/algorithms"

func WaysOfAnagrams(source, target string) int {
	// Total ways we could form target string from subset of source string
	if len(source) < len(target) {
		return 0
	}

	sourceCount := map[rune]int{}
	targetCount := map[rune]int{}
	for _, ch := range source {
		sourceCount[ch]++
	}
	for _, ch := range target {
		targetCount[ch]++
	}

	ways := 1
	for ch, freq := range targetCount {
		if sourceCount[ch] < freq {
			return 0
		}

		ways *= algorithms.Combination(sourceCount[ch], freq)
	}

	return ways
}
