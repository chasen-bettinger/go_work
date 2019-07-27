package strain

// import "log"

// Ints is a map of integers
type Ints []int

// Lists is a map of a map of integers
type Lists [][]int

// Strings is a map of strings
type Strings []string

// Keep takes integers and returns the set of them that match the function evaluation
func (integersIn Ints) Keep(f func(int) bool) Ints {
	if integersIn == nil {
		return nil
	}
	integersOut := Ints{}
	for _, val := range integersIn {
		if f(val) {
			integersOut = append(integersOut, val)
		}
	}
	return integersOut
}

// Discard takes integers and returns the set of them that don't match the function evaluation
func (integersIn Ints) Discard(f func(int) bool) Ints {
	if integersIn == nil {
		return nil
	}
	integersOut := Ints{}
	for _, val := range integersIn {
		if !f(val) {
			integersOut = append(integersOut, val)
		}
	}
	return integersOut
}

// Keep takes a list of integers and returns the set of them that match the function evaluation
func (listIn Lists) Keep(f func([]int) bool) Lists {
	if listIn == nil {
		return nil
	}
	listOut := Lists{}
	for _, val := range listIn {
		if f(val) {
			listOut = append(listOut, val)
		}
	}
	return listOut
}

// Keep takes strings and returns the set of them that match the function evaluation
func (stringIn Strings) Keep(f func(string) bool) Strings {
	if stringIn == nil {
		return nil
	}
	stringOut := Strings{}
	for _, val := range stringIn {
		if f(val) {
			stringOut = append(stringOut, val)
		}
	}
	return stringOut
}
