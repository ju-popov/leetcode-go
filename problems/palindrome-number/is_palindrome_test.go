package palindromenumber_test

import (
	"reflect"
	"testing"

	palindromenumber "github.com/ju-popov/leetcode-go/problems/palindrome-number"
)

func IsPalindrome(x int) bool {
	return false
}

type testCase struct {
	name     string
	x        int
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "01",
		x:        121,
		expected: true,
	},
	{
		name:     "-121",
		x:        -121,
		expected: false,
	},
	{
		name:     "03",
		x:        10,
		expected: false,
	},
}

func TestIsPalindrome(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		currentTestCase := testCase
		t.Run(currentTestCase.name, func(t *testing.T) {
			t.Parallel()
			actual := palindromenumber.IsPalindrome(currentTestCase.x)
			if !reflect.DeepEqual(actual, currentTestCase.expected) {
				t.Errorf("Expected result for test: '%v' is: '%v', but the actual result is: '%v'", currentTestCase.name, currentTestCase.expected, actual)
			}
		})
	}
}
