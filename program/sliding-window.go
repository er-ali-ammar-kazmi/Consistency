package program

func LengthOflongestSubstring(input string) string {
	length := 0
	substring := ""
	seen := map[rune]int{}
	start := 0
	for idx, chr := range input {
		if seen[chr] != 0 {
			start = seen[chr]
		}
		seen[chr] = idx + 1
		temp := input[start : idx+1]
		if len(temp) > length {
			substring = temp
			length = len(temp)
		}
	}
	return substring
}
