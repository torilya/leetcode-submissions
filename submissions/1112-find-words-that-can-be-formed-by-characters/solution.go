func countCharacters(words []string, chars string) int {
	alphabetLen := int('z' - 'a' + 1)
	bytesCountChars := make([]int, alphabetLen)

    for i := 0; i < len(chars); i++ {
        bytesCountChars[chars[i]-'a']++
    }

    var goodStrsLensSum int

    for _, word := range words {
        wordLen := len(word)
        bytesCountWord := make([]int, alphabetLen)

        for i := 0; i < wordLen; i++ {
            bytesCountWord[word[i]-'a']++
        }

        isWordGood := true

        for i := 0; i < alphabetLen; i++ {
            if bytesCountWord[i] > bytesCountChars[i] {
                isWordGood = false
                break
            }
        }

        if isWordGood {
            goodStrsLensSum += wordLen
        }
    }

    return goodStrsLensSum
}

