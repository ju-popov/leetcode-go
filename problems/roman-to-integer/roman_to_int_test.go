package romantointeger_test

import (
	"reflect"
	"testing"

	romantointeger "github.com/ju-popov/leetcode-go/problems/roman-to-integer"
)

type testCase struct {
	name     string
	s        string
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "01",
		s:        "III",
		expected: 3,
	},
	{
		name:     "02",
		s:        "LVIII",
		expected: 58,
	},
	{
		name:     "03",
		s:        "MCMXCIV",
		expected: 1994,
	},
}

func TestRomanToInt(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		currentTestCase := testCase
		t.Run(currentTestCase.name, func(t *testing.T) {
			t.Parallel()
			actual := romantointeger.RomanToInt(currentTestCase.s)
			if !reflect.DeepEqual(actual, currentTestCase.expected) {
				t.Errorf("Expected result for test: '%v' is: '%v', but the actual result is: '%v'", currentTestCase.name, currentTestCase.expected, actual)
			}
		})
	}
}
