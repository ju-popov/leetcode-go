package palindromenumber

func IsPalindrome(x int) bool { //nolint:varnamelen
	if x < 0 {
		return false
	}

	if x >= 0 && x <= 9 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	xOrig := x
	xPalindrome := 0

	for x > 0 {
		xPalindrome = xPalindrome*10 + x%10 //nolint:gomnd
		x /= 10
	}

	return xOrig == xPalindrome
}
