package accumulate

// Accumulate takes input and a function and applies that function to the input
func Accumulate(input []string, accumlator func(string) string) []string {
	res := make([]string, len(input))
	for key, val := range input {
		res[key] = accumlator(val)
	}
	return res
}