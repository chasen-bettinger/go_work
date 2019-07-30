package combinationsum

import (
	"reflect"
	"testing"
)

type combinationTests struct {
	candidates []int
	target     int
	expected   [][]int
}

var combinationTestCases = []combinationTests{
	{
		candidates: []int{2, 3, 6, 7},
		target:     7,
		expected:   [][]int{{2, 2, 3}, {7}},
	},
	{
		candidates: []int{8, 7, 4, 3},
		target:     11,
		expected:   [][]int{{3, 4, 4}, {3, 8}, {4, 7}},
	},
}

func TestCounts(t *testing.T) {
	for _, v := range combinationTestCases {
		combo := combinationSum(v.candidates, v.target)
		if reflect.DeepEqual(combo, v.expected) {
			t.Logf("PASS")
		} else {
			t.Fatalf("FAIL: Expected: %#v\"", v.expected)
		}
	}
}
