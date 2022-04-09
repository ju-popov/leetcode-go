# 20. Valid Parentheses

**Difficulty**: Easy

**Source**: https://leetcode.com/problems/valid-parentheses/

## Related Topics:

* [String](https://leetcode.com/tag/string/)
* [Stack](https://leetcode.com/tag/stack/)

## Problem:

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

**Example 1:**

```
Input: s = "()"
Output: true
```

**Example 2:**

```
Input: s = "()[]{}"
Output: true
```

**Example 3:**

```
Input: s = "(]"
Output: false
```

**Constraints:**

- `1 <= s.length <= 104`
- `s` consists of parentheses only `'()[]{}'`.

## Solution:

```go
package validparentheses

func IsValid(s string) bool {
	parentheses := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []int32

	for _, char := range s {
		if value, ok := parentheses[char]; ok {
			stack = append(stack, value)
		} else if len(stack) == 0 {
			return false
		} else if char != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
```