func canPlaceFlowers(flowerbed []int, n int) bool {
    if n <= 0 {
        return true
    }
    
    for i, plot := range flowerbed {
        if plot == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
            flowerbed[i] = 1
            n--

            if n <= 0 {
                return true
            }
        }
    }

    return false
}

