package validparentheses_test

import (
	"reflect"
	"testing"

	validparentheses "github.com/ju-popov/leetcode-go/problems/valid-parentheses"
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
		s:        "()",
		expected: true,
	},
	{
		name:     "02",
		s:        "()[]{}",
		expected: true,
	},
	{
		name:     "03",
		s:        "(]",
		expected: false,
	},
}

func TestIsValid(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		currentTestCase := testCase
		t.Run(currentTestCase.name, func(t *testing.T) {
			t.Parallel()
			actual := validparentheses.IsValid(currentTestCase.s)
			if !reflect.DeepEqual(actual, currentTestCase.expected) {
				t.Errorf("Expected result for test: '%v' is: '%v', but the actual result is: '%v'", currentTestCase.name, currentTestCase.expected, actual)
			}
		})
	}
}
