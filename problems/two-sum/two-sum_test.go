package twosum_test

import (
	"reflect"
	"testing"

	twosum "github.com/ju-popov/leetcode-go/problems/two-sum"
)

type testCase struct {
	name     string
	nums     []int
	target   int
	expected []int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "01",
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		name:     "02",
		nums:     []int{3, 2, 4, 6},
		target:   6,
		expected: []int{1, 2},
	},
	{
		name:     "03",
		nums:     []int{3, 3, 6},
		target:   6,
		expected: []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		currentTestCase := testCase
		t.Run(currentTestCase.name, func(t *testing.T) {
			t.Parallel()
			actual := twosum.TwoSum(currentTestCase.nums, currentTestCase.target)
			if !reflect.DeepEqual(actual, currentTestCase.expected) {
				t.Errorf("Expected result for test: '%v' is: '%v', but the actual result is: '%v'", currentTestCase.name, currentTestCase.expected, actual)
			}
		})
	}
}
