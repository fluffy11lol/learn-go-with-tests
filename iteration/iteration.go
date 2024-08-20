package iteration

import "strings"

func Repeat1(character string, count int) string {
	res := make([]string, count)
	for i := 0; i < count; i++ {
		res[i] = character
	}
	return strings.Join(res, "")
}
func Repeat2(character string, count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res += character
	}
	return res
}

// Repeat3 The fastest func
func Repeat3(character string, count int) string {
	res := strings.Repeat(character, count)
	return res
}

func Repeat4(character string, count int) string {
	if count <= 0 {
		return ""
	}
	var builder strings.Builder
	builder.Grow(count * len(character))

	for i := 0; i < count; i++ {
		builder.WriteString(character)
	}

	return builder.String()
}
