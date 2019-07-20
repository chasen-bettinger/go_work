// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(val int) bool {
	var evenBy4 = val % 4 == 0 
	if (!evenBy4) { 
		return false
	}

	var evenBy100 = val % 100 == 0
	var evenBy400 = val % 400 == 0
	var evenBy4and100 = evenBy4 && evenBy100
	var allEven = evenBy4 && evenBy100 && evenBy400

	if (evenBy4and100 && !allEven) {
		return false
	} 
	return true
}
