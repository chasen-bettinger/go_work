package one

import (
	"math"
)

func FuelAlgo(number int) int {
	return int(math.Floor(float64(number)/3) - 2)
}

// GetFuel should have a comment documenting it.
func GetFuel(modules []int) int {
	fuel := 0
	for _, module := range modules {
		fuelOfModule := FuelAlgo(module)
		fuel += fuelOfModule

		for fuelOfModule > 0 {
			fuelOfModule = FuelAlgo(fuelOfModule)
			if fuelOfModule > 0 {
				fuel += fuelOfModule
			}
		}
	}

	return fuel
}
