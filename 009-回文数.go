package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var revertedNumber int
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x = x/10
	}
	if x == revertedNumber || x == revertedNumber/10 {
		return true
	}
	return false
}
