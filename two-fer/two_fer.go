package twofer

import (
	"fmt"
)

// ShareWith prints a statement with a specific name if given
func ShareWith(name string) string {
	var printedName = "you"
	var nameIsNotNull = len(name) > 0
	if nameIsNotNull {
		printedName = name
	}
	return fmt.Sprintf("One for %v, one for me.", printedName)
}
