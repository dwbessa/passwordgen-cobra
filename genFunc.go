package main

func generateRandomString(length int, charSet []rune) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = charSet[rng.Intn(len(charSet))] // Use rng.Intn instead of rand.Intn
	}
	return string(b)
}

func lowercaseRunes(s []rune) []rune {
	result := make([]rune, len(s))
	for i, r := range s {
		if 'A' <= r && r <= 'Z' {
			r += 'a' - 'A' // Convert to lowercase
		}
		result[i] = r
	}
	return result
}
