package hamming

import (
	"errors"
	"strings"
)

// Distance calculates the hamming distance between two strands of DNA
func Distance(a, b string) (int, error) {

	if a == "" && b == "" {
		return 0, nil
	}

	if len(a) != len(b) {
		return -1, errors.New("strands must be of the same size")
	}

	splitA := strings.Split(a, "")
	splitB := strings.Split(b, "")
	hammingDistance := 0

	for key := range splitA {
		nucleotideA := splitA[key]
		nucleotideB := splitB[key]

		if nucleotideA != nucleotideB {
			hammingDistance++
		}

	}

	return hammingDistance, nil
}
