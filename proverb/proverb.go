// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverb := []string{}
	lengthOfRhyme := len(rhyme)

	for key, val := range rhyme {
		if key < lengthOfRhyme - 1 {
			proverb = proverb.append(proverb, Line(val, rhyme[key+1]))
		}
	}
	return []string{}
}

func Line(want, lost string) string {
	return "For want of a " + want + " the  " + lost + " was lost."
} 
