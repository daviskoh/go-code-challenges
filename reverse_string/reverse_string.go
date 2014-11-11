package reverse_string

func ReverseString(s string) string {
	if len(s) < 2 {
		return s
	}

	return s[len(s)-1:len(s)] + ReverseString(s[0:len(s)-1])
}
