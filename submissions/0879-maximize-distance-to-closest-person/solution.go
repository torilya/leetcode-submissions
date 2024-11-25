func maxDistToClosest(seats []int) int {
    seatsLen := len(seats)
	var anchor int
	var distToClosestMax int

	for i := 1; i < seatsLen; i++ {
		if seats[i-1] > seats[i] {
			anchor = i
		}

		if seats[i-1] < seats[i] {
			if anchor == 0 {
				distToClosestMax = i
			} else {
				distToClosestMax = max(distToClosestMax, (len(seats[anchor:i])+1)/2)
			}
		}

		if i == seatsLen-1 && seats[i] == 0 {
			distToClosestMax = max(distToClosestMax, len(seats[anchor:]))
		}
	}

	return distToClosestMax
}

