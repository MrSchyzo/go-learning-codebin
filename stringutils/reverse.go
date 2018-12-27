package stringutils

// ReverseString is a simple function
func ReverseString(s string) string {

	r := []rune(s)
	l := len(r)
	last := l - 1

	for i := 0; i < l/2; i = i + 1 {
		r[i], r[last-i] = r[last-i], r[i]
	}

	return string(r)

}
