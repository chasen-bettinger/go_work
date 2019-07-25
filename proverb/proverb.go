package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	lengthOfRhyme := len(rhyme)
	if lengthOfRhyme == 0 {
		return rhyme 
	}
	proverb := make([]string, lengthOfRhyme)

	for key, val := range rhyme {
		if key < lengthOfRhyme - 1 {
			want := val
			lost := rhyme[key+1]
			proverb[key] = fmt.Sprintf("For want of a %v the %v was lost.", want, lost)
		}
	}
	proverb[lengthOfRhyme - 1] = fmt.Sprintf("And all for the want of a %v.", rhyme[0])
	return proverb
}
