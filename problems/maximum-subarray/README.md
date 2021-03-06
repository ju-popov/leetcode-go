# 53. Maximum Subarray

**Difficulty**: Easy

**Source**: https://leetcode.com/problems/maximum-subarray/

## Related Topics:

* [Array](https://leetcode.com/tag/array/)
* [Divide and Conquer](https://leetcode.com/tag/divide-and-conquer/)
* [Dynamic Programming](https://leetcode.com/tag/dynamic-programming/)

## Problem:

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return *its sum*.

A **subarray** is a **contiguous** part of an array.

**Example 1:**

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

**Example 2:**

```
Input: nums = [1]
Output: 1
```

**Example 3:**

```
Input: nums = [5,4,-1,7,8]
Output: 23
```

**Constraints:**

- `1 <= nums.length <= 105`
- `-104 <= nums[i] <= 104`

**Follow up:** If you have figured out the `O(n)` solution, try coding another solution using the **divide and conquer** approach, which is more subtle.

## Solution:

```go
package maximumsubarray

func MaxSubArray(nums []int) int {
	var (
		sum    int
		maxSum int
	)

	for index, num := range nums {
		if index == 0 {
			sum = num
			maxSum = num

			continue
		}

		if sum < 0 {
			sum = 0
		}

		sum += num

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
```

