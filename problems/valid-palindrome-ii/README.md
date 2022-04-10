# 680. Valid Palindrome II

**Difficulty**: Easy

**Source**: https://leetcode.com/problems/valid-palindrome-ii/

## Related Topics:

* [Two Pointers](https://leetcode.com/tag/two-pointers/)
* [String](https://leetcode.com/tag/string/)
* [Greedy](https://leetcode.com/tag/greedy/)

## Problem:

Given a string `s`, return `true` *if the*`s`*can be palindrome after deleting **at most one** character from it*.

**Example 1:**

```
Input: s = "aba"
Output: true
```

**Example 2:**

```
Input: s = "abca"
Output: true
Explanation: You could delete the character 'c'.
```

**Example 3:**

```
Input: s = "abc"
Output: false
```

**Constraints:**

- `1 <= s.length <= 105`
- `s` consists of lowercase English letters.

## Solution:

```go
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
```