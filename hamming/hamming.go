package hamming

import (
	"errors"
)

// Distance calculates the hamming distance between two strands of DNA
func Distance(a, b string) (int, error) {

	hammingDistance := 0

	if len(a) != len(b) {
		return -1, errors.New("strands must be of the same size")
	}

	for key := range a {
		if a[key] != b[key] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
