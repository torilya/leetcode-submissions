import 
(
    "fmt"
    "strings"
)

func reverseWords(s string) string {
    words := strings.Split(s, " ")
    var reversedWordsStr string

    for i, word := range words {
        runes := []rune(word)

        for j := len(runes)-1; j >= 0; j-- {
            reversedWordsStr = fmt.Sprintf("%s%c", reversedWordsStr, runes[j])
        }

        if i < len(words)-1 {
            reversedWordsStr = fmt.Sprintf("%s ", reversedWordsStr)
        }
    }

    return reversedWordsStr
}

