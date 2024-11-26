func maxDistToClosest(seats []int) int {
	prev := -1
	var distToClosestMax int

	for i, seat := range seats {
		if seat == 1 {
			if prev == -1 {
				distToClosestMax = i
			} else {
				distToClosestMax = max(distToClosestMax, (i-prev)/2)
			}

            prev = i
		}
	}

	return max(distToClosestMax, len(seats)-prev-1)
}

