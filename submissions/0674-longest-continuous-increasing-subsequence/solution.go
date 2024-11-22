func findLengthOfLCIS(nums []int) int {
	var lcisLen int
	var anchor int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] >= nums[i] {
			anchor = i
		}
        
		lcisLen = max(lcisLen, i-anchor+1)
	}

	return lcisLen
}

