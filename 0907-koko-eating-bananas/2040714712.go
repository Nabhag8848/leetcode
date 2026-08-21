func minEatingSpeed(piles []int, h int) int {
    minPossibleHour := 1
    maxPossibleHour := findMaximum(piles)

    for minPossibleHour <=  maxPossibleHour {
        hourly := minPossibleHour + (maxPossibleHour - minPossibleHour) / 2
        hour := calculateTotalHours(piles, hourly)

        if hour < h {
            maxPossibleHour = hourly - 1
        } else if hour > h {
            minPossibleHour = hourly + 1
        } else {
            maxPossibleHour = hourly - 1
        }
    }

    return minPossibleHour
}

func calculateTotalHours(piles []int, hourly int) int {
    totalHours := 0

    for i := range piles {
        totalHours += int(math.Ceil(float64(piles[i]) / float64(hourly)))
    }

    return totalHours
}

func findMaximum(piles []int) int {
    max := math.MinInt32

    for i := range piles {
        if max < piles[i] {
            max = piles[i]
        }
    }

    return max
}

/*
    4 11 20 23 30 
*/