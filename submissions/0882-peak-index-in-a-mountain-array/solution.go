func peakIndexInMountainArray(arr []int) int {
	leftIndex := 1
	rightIndex := len(arr) - 2

	for leftIndex < rightIndex {
		midIndex := (leftIndex + rightIndex) / 2

		if arr[midIndex] < arr[midIndex+1] {
			leftIndex = midIndex + 1
		} else {
			rightIndex = midIndex
		}
	}

	return leftIndex
}

