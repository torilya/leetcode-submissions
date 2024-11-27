func romanToInt(s string) int {
	runes := []rune(s)

	romanNumeralsInts := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var integer int

	for i := len(runes)-1; i >= 0; i-- {
		if i < len(runes)-1 && romanNumeralsInts[runes[i]] < romanNumeralsInts[runes[i+1]] {
			integer -= romanNumeralsInts[runes[i]]
            continue
		}

        integer += romanNumeralsInts[runes[i]]
	}

	return integer
}

