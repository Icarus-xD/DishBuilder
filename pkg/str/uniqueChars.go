package str

func UniqueChars(s string) []rune {
	charMap := make(map[rune]bool)
	uniqueChars := []rune{}

	for _, char := range s {
		if _, ok := charMap[char]; !ok {
			charMap[char] = true
			uniqueChars = append(uniqueChars, char)
		}
	}

	return uniqueChars
}