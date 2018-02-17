package reverse

// String reverses the given string
func String(str string) string {
	var reversed []byte
	for i := len(str) - 1; i >= 0; i-- {
		reversed = append(reversed, str[i])
	}
	return string(reversed)
}
