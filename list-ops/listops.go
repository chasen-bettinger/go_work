package listops

// IntList is a list of integers
type IntList []int

// binFunc is a function that takes integers and returns an integer
type binFunc func(x, y int) int

// predFunc is a function that takes an integer, and returns a boolean
type predFunc func(n int) bool

// predFunc is a function that takes an integer, and returns an integer
type unaryFunc func(n int) int

// Foldr takes an acculumator and applies that against a list from right to left
func (intList IntList) Foldr(fn binFunc, initial int) (out int) {
	i := len(intList) - 1
	out = initial
	for range intList {
		v := intList[i]
		out = fn(v, out)
		i--
	}
	return
}

// Foldl takes an acculumator and applies that against a list from left to right
func (intList IntList) Foldl(fn binFunc, initial int) (out int) {
	out = initial
	for _, v := range intList {
		out = fn(out, v)
	}
	return
}

// Filter takes a function that is applied against the list and returns the items that meet the conditions in the function
func (intList IntList) Filter(fn predFunc) (out IntList) {
	out = IntList{}
	for _, v := range intList {
		if fn(v) {
			out = append(out, v)
		}
	}
	return
}

// Length provides the length of a list of integers
func (intList IntList) Length() (out int) {
	out = 0
	for range intList {
		out++
	}
	return
}

// Map goes over a series of integers and applies a function to them
func (intList IntList) Map(fn unaryFunc) (out IntList) {
	out = IntList{}
	for _, v := range intList {
		out = append(out, fn(v))
	}
	return
}

// Reverse takes a list and reverses the elements in that list
func (intList IntList) Reverse() (out IntList) {
	out = IntList{}
	i := len(intList) - 1
	for range intList {
		out = append(out, intList[i])
		i--
	}
	return
}

// Append takes a list and appends that to the list that called this function
func (intList IntList) Append(newIntList IntList) (out IntList) {
	out = intList
	for _, v := range newIntList {
		out = append(out, v)
	}
	return
}

// Concat concatenates two arrays
func (intList IntList) Concat(newIntList []IntList) (out IntList) {
	out = intList
	for _, v := range newIntList {
		for _, k := range v {
			out = append(out, k)
		}
	}
	return
}
