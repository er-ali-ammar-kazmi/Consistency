package program

func IsStrPalindrome(input string) bool {
	var reverse string
	for _, str := range input {
		reverse = string(str) + reverse
	}
	if input == reverse {
		return true
	}
	return false
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
