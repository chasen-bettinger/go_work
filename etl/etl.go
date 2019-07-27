package etl

import "strings"

// Transform performs the T in ETL
func Transform(input map[int][]string) map[string]int {
	var transformedResult = make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			transformedResult[strings.ToLower(letter)] = score
		}
	}
	return transformedResult
}