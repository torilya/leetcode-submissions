func romanToInt(s string) int {
	romanNumeralsInts := [256]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	integer := romanNumeralsInts[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		numCurrent := romanNumeralsInts[s[i]]

		if numCurrent < romanNumeralsInts[s[i+1]] {
			integer -= numCurrent
            continue
		}
        
        integer += numCurrent
	}

	return integer
}

