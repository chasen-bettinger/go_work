package collatzconjecture

import "errors"

// CollatzConjecture gets the number of steps it took to complete the CollatzConjecture
func CollatzConjecture(number int) (int, error) {

	if number < 1 {
		return -1, errors.New("number must be a positive number")
	}

	steps := 0

	if number == 1 {
		return steps, nil
	}

	var performArithmetic func()

	performArithmetic = func() {
		numberIsEven := number%2 == 0
		if numberIsEven {
			number = number / 2
		} else {
			number = number*3 + 1
		}
		steps++
		if number != 1 {
			performArithmetic()
		}
	}

	performArithmetic()
	return steps, nil
}
