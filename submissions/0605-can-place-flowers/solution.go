func canPlaceFlowers(flowerbed []int, n int) bool {
    if n <= 0 {
        return true
    }
    
    flowerbedLen := len(flowerbed)

    for i := 0; i < flowerbedLen; i++ {
        if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == flowerbedLen-1 || flowerbed[i+1] == 0) {
            flowerbed[i] = 1
            n--

            if n <= 0 {
                return true
            }

            i++
        }
    }

    return false
}

