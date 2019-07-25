package raindrops

import "strconv"

// Convert converts a number of raindrops to its sound based on factors of that number
func Convert(raindrops int) string {

	message := ""

	if raindrops%3 == 0 {
		message += "Pling"
	}

	if raindrops%5 == 0 {
		message += "Plang"
	}

	if raindrops%7 == 0 {
		message += "Plong"
	}

	if message == "" {
		return strconv.Itoa(raindrops)
	}

	return message
}
