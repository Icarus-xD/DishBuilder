package combination

func factorial(number int) int {

	if number == 1 {
		return 1
	}

	factorialOfNumber := number * factorial(number-1)

	return factorialOfNumber
}

func CombinationCount(rows, count int) int {
	return factorial(rows) / (factorial(count) * factorial(rows-count))
}