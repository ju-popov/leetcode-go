package validpalindromeii_test

import (
	"reflect"
	"testing"

	validpalindromeii "github.com/ju-popov/leetcode-go/problems/valid-palindrome-ii"
)

type testCase struct {
	name     string
	s        string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "01",
		s:        "aba",
		expected: true,
	},
	{
		name:     "02",
		s:        "abca",
		expected: true,
	},
	{
		name:     "03",
		s:        "abc",
		expected: false,
	},
}

func TestTwoSum(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		currentTestCase := testCase
		t.Run(currentTestCase.name, func(t *testing.T) {
			t.Parallel()
			actual := validpalindromeii.ValidPalindrome(currentTestCase.s)
			if !reflect.DeepEqual(actual, currentTestCase.expected) {
				t.Errorf("Expected result for test: '%v' is: '%v', but the actual result is: '%v'", currentTestCase.name, currentTestCase.expected, actual)
			}
		})
	}
}
