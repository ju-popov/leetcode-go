package maximumsubarray_test

import (
	"reflect"
	"testing"

	maximumsubarray "github.com/ju-popov/leetcode-go/problems/maximum-subarray"
)

type testCase struct {
	name     string
	nums     []int
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "01",
		nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 6,
	},
	{
		name:     "02",
		nums:     []int{1},
		expected: 1,
	},
	{
		name:     "03",
		nums:     []int{5, 4, -1, 7, 8},
		expected: 23,
	},
}

func TestMaxSubArray(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		currentTestCase := testCase
		t.Run(currentTestCase.name, func(t *testing.T) {
			t.Parallel()
			actual := maximumsubarray.MaxSubArray(currentTestCase.nums)
			if !reflect.DeepEqual(actual, currentTestCase.expected) {
				t.Errorf("Expected result for test: '%v' is: '%v', but the actual result is: '%v'", currentTestCase.name, currentTestCase.expected, actual)
			}
		})
	}
}
