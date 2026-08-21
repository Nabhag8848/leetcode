func maxDistance(position []int, m int) int {
    sort.Ints(position)
    low := 1
    high := position[len(position) - 1]

    for low <= high {
        mid := low + (high - low) / 2

        if canbePlaced(position, mid, m) {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return high
}   

func canbePlaced(position []int, distance int, balls int) bool {
    countBalls := 1 
    last_pos := position[0]

    for i:=1; i < len(position); i++ {
        if position[i] - last_pos >= distance {
            countBalls++ 
            last_pos = position[i]
        }

        if countBalls >= balls {
            return true
        }
    }

    return false
}