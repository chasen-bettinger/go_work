package space

import (
	"math"
)

// Planet a string that represents a planet
type Planet string

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func roundBy2(num float64) float64 {
	output := math.Pow(10, float64(2))
	return float64(round(num * output)) / output
}

// Age given an age in seconds, calculate hold someone would be on
func Age(seconds float64, planet Planet) float64 {
	var oneEarthYearInSeconds = float64(31557600)
	var earthYears = float64(seconds / oneEarthYearInSeconds)
	var roundedYears = roundBy2(earthYears)
	if (planet == "Earth") {
		return roundedYears
	}
	planetsInEarthYears := map[Planet]float64 {
		"Mercury": 0.2408467,
		"Venus": 0.61519726,
		"Mars": 1.8808158,
		"Jupiter": 11.862615,
		"Saturn": 29.447498,
		"Uranus": 84.016846,
		"Neptune": 164.79132,
	}
	return roundBy2(roundedYears / planetsInEarthYears[planet])
}

