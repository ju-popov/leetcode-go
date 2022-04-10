package longestcommonprefix_test

import (
	"reflect"
	"testing"

	longestcommonprefix "github.com/ju-popov/leetcode-go/problems/longest-common-prefix"
)

type testCase struct {
	name     string
	strs     []string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "01",
		strs:     []string{"flower", "flow", "flight"},
		expected: "fl",
	},
	{
		name:     "02",
		strs:     []string{"dog", "racecar", "car"},
		expected: "",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		currentTestCase := testCase
		t.Run(currentTestCase.name, func(t *testing.T) {
			t.Parallel()
			actual := longestcommonprefix.LongestCommonPrefix(currentTestCase.strs)
			if !reflect.DeepEqual(actual, currentTestCase.expected) {
				t.Errorf("Expected result for test: '%v' is: '%v', but the actual result is: '%v'", currentTestCase.name, currentTestCase.expected, actual)
			}
		})
	}
}
