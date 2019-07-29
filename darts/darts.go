package darts

import "math"

// Score takes coordiates and gives a score based on where those coordinates are on a dart board that has a radius of 10
func Score(x, y float64) (score int) {
	coords := math.Pow(y, 2) + math.Pow(x, 2)
	score = 0

	if coords <= 1e2 {
		score = 1
	}

	if coords <= 25 {
		score = 5
	}

	if coords <= 1 {
		score = 10
	}

	return
}
