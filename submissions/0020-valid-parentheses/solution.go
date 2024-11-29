func isValid(s string) bool {
    if len(s)%2 == 1 {
        return false
    }

    parensOpenClose := map[byte]byte{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    var parensOpenStack []byte

    for i := 0; i < len(s); i++ {
        if _, ok := parensOpenClose[s[i]]; ok {
            parensOpenStack = append(parensOpenStack, s[i])

            continue
        }

        if len(parensOpenStack) == 0 || s[i] != parensOpenClose[parensOpenStack[len(parensOpenStack)-1]] {
            return false
        }

        parensOpenStack = parensOpenStack[:len(parensOpenStack)-1]
    }

    return len(parensOpenStack) == 0
}

