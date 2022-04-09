package besttimetobuyandsellstock_test

import (
	"reflect"
	"testing"

	besttimetobuyandsellstock "github.com/ju-popov/leetcode-go/problems/best-time-to-buy-and-sell-stock"
)

type testCase struct {
	name     string
	prices   []int
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "01",
		prices:   []int{7, 1, 5, 3, 6, 4},
		expected: 5,
	},
	{
		name:     "02",
		prices:   []int{7, 6, 4, 3, 1},
		expected: 0,
	},
}

func TestMaxProfit(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		currentTestCase := testCase
		t.Run(currentTestCase.name, func(t *testing.T) {
			t.Parallel()
			actual := besttimetobuyandsellstock.MaxProfit(currentTestCase.prices)
			if !reflect.DeepEqual(actual, currentTestCase.expected) {
				t.Errorf("Expected result for test: '%v' is: '%v', but the actual result is: '%v'", currentTestCase.name, currentTestCase.expected, actual)
			}
		})
	}
}
