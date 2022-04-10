# 9. Palindrome Number

**Difficulty**: Easy

**Source**: https://leetcode.com/problems/palindrome-number/

## Related Topics:

* [Math](https://leetcode.com/tag/math/)

## Problem:

Given an integer `x`, return `true` if `x` is palindrome integer.

An integer is a **palindrome** when it reads the same backward as forward.

- For example, `121` is a palindrome while `123` is not.

**Example 1:**

```
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
```

**Example 2:**

```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

**Example 3:**

```
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

**Constraints:**

- `-231 <= x <= 231 - 1`

**Follow up:** Could you solve it without converting the integer to a string?

## Solution:

```go
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
```