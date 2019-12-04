package two

import (
	"fmt"
)

// CodeTernary should have a comment documenting it.
func CodeTernary(firstCode int, a int, b int) int {
	if firstCode == 1 {
		return a
	}

	return b
}

// RestoreGravityAssistProgram should have a comment documenting it.
func RestoreGravityAssistProgram(codes []int, noun int, verb int) int {
	fmt.Println(noun, verb)
	codes = append([]int(nil), codes...)
	codes[1] = noun
	codes[2] = verb
	// codes[1] = 12
	// codes[2] = 2
	// if noun != 0 {
	// }

	// if verb != 0 {
	// }

	currentPosition := 0
	for {
		fmt.Println(codes[currentPosition], codes[currentPosition+1], codes[currentPosition+2], codes[currentPosition+3])
		firstCode := codes[currentPosition]
		fmt.Println(firstCode)
		if firstCode == 99 {
			break
		}

		indexOfOne := codes[currentPosition+1]
		indexOfTwo := codes[currentPosition+2]
		indexOfThree := codes[currentPosition+3]
		addition := codes[indexOfOne] + codes[indexOfTwo]
		multiplication := codes[indexOfOne] * codes[indexOfTwo]

		newValue := CodeTernary(firstCode, addition, multiplication)
		codes[indexOfThree] = newValue
		currentPosition = currentPosition + 4
	}
	return codes[0]
}

// GetNounVerb should have a comment documenting it.
func GetNounVerb(codes []int, target int) int {
	noun := 0
	verb := 99

	for {
		result := RestoreGravityAssistProgram(codes, noun, verb)
		if result == target {
			break
		}

		// do binary search

		if result > target {
			verb = verb - 1
		} else {
			noun = noun + 1
		}
	}

	return 100*noun + verb
}
