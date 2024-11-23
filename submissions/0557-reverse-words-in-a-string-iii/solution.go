func reverse(bytes []byte) {
	bytesLen := len(bytes)

	for i := 0; i < bytesLen/2; i++ {
		bytes[i], bytes[bytesLen-1-i] = bytes[bytesLen-1-i], bytes[i]
	}
}

func reverseWords(s string) string {
	bytes := []byte(s)
	wordStartIndex := 0

	for i, r := range bytes {
		if r == ' ' {
			reverse(bytes[wordStartIndex:i])

            wordStartIndex = i + 1
		}
	}

	reverse(bytes[wordStartIndex : len(bytes)])

	return string(bytes)
}

