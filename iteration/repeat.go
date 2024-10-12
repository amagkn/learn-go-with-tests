package iteration

const repeatCount = 5

// Repeat char "count" times
func Repeat(character string, count int) string {
	var repeated string

	if count == 0 {
		count = repeatCount
	}

	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
