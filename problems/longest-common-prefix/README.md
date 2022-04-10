# 14. Longest Common Prefix

**Difficulty**: Easy

**Source**: https://leetcode.com/problems/longest-common-prefix/

##  Related Topics:

* [String](https://leetcode.com/tag/string/)

## Problem:

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string `""`.

**Example 1:**

```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

**Example 2:**

```
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

**Constraints:**

- `1 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` consists of only lower-case English letters.

## Solution:

```go
package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	var minLen int

	for i, str := range strs {
		l := len(str)
		if i == 0 {
			minLen = l
		} else if l < minLen {
			minLen = l
		}
	}

	for pos := 0; pos < minLen; pos++ {
		var char uint8

		for i, str := range strs {
			if i == 0 {
				char = str[pos]
			} else if char != str[pos] {
				return strs[0][:pos]
			}
		}
	}

	return strs[0][:minLen]
}
```