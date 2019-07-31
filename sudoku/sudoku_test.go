package sudoku

import (
	"reflect"
	"testing"
)

type sudokuTests struct {
	board    [][]byte
	expected bool
}

var sudokuTestCases = []sudokuTests{
	{
		board: [][]byte{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		expected: false,
	},
}

func TestCounts(t *testing.T) {
	for _, v := range sudokuTestCases {
		combo := isValidSudoku(v.board)
		if reflect.DeepEqual(combo, v.expected) {
			t.Logf("PASS")
		} else {
			t.Fatalf("FAIL: Expected: %#v\"", v.expected)
		}
	}
}
