// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear evaluates whether a given year is a leap year
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
