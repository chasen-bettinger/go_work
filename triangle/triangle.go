package triangle

import "math"

// Kind type of triangle we are returning
type Kind string

const (
	// NaT not a triangle
	NaT = "NaT"
	// Equ equilateral
	Equ = "Equ"
	// Iso isosceles
	Iso = "Iso"
	// Sca scalene
	Sca = "Sca"
)

// KindFromSides detemines the type of triangle based on the length of sides passed into it
func KindFromSides(a, b, c float64) Kind {
	if !isATriangle(a, b, c) {
		return NaT
	}

	var lengths map[float64]float64
	lengths = make(map[float64]float64)

	lengths[a]++
	lengths[b]++
	lengths[c]++

	for _, length := range lengths {
		if length == 3 {
			return Equ
		}
		if length == 2 {
			return Iso
		}
	}

	return Sca
}

func isATriangle(a, b, c float64) bool {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a+b < c || a+c < b || c+b < a {
		return false
	}

	return true

}
