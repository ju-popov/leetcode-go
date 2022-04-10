package validpalindromeii

func validPalindromeHelper(s string, deletions int) bool { //nolint:varnamelen
	if len(s) <= 1 {
		return true
	}

	if s[0] == s[len(s)-1] {
		return validPalindromeHelper(s[1:len(s)-1], deletions)
	}

	if deletions == 0 {
		return false
	}

	if validPalindromeHelper(s[:len(s)-1], deletions-1) {
		return true
	}

	if validPalindromeHelper(s[1:], deletions-1) {
		return true
	}

	return false
}

func ValidPalindrome(s string) bool {
	return validPalindromeHelper(s, 1)
}
