package twofer

import (
	"fmt"
)

// ShareWith prints a statement with a specific name if given
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
