package combination

func GetCombinations[T any](rows []T, count int) [][]T {
	var result [][]T
	combination := make([]T, count)
	generateCombinations(rows, count, 0, combination, &result)
	return result
}

func generateCombinations[T any](
	rows []T, count int, index int,
	combination []T, result *[][]T,
) {
	if index == count {
		c := make([]T, count)
		copy(c, combination)
		*result = append(*result, c)
		return
	}

	for i := 0; i <= len(rows) - count + index; i++ {
		combination[index] = rows[i]
		generateCombinations(rows[i + 1:], count, index + 1, combination, result)
	}
}