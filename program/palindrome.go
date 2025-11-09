package program

func IsStrPalindrome(input string) bool {
	endPointer := len(input) - 1
	flag := true
	if endPointer < 0 {
		return false
	}
	for idx, str := range input {
		if endPointer <= idx {
			break
		}
		if str == rune(input[endPointer]) {
			endPointer--
			continue
		} else {
			flag = false
			break
		}
	}
	return flag
}

func IsIntPalindrome(input int) bool {

	var reverse int
	var tmp = input
	for tmp > 0 {
		rem := tmp % 10
		reverse = reverse*10 + rem
		tmp = tmp / 10
	}
	return input == reverse
}
