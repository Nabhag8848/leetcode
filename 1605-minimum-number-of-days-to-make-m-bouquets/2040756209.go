func minDays(bloomDay []int, m int, k int) int {
    numberOfFlowers := len(bloomDay)
    numberOfBouquets := m
    numberOfFlowersIneach := k

    if numberOfBouquets * numberOfFlowersIneach > numberOfFlowers { 
        return -1
    }

    minPossibleDays := findMinimumNumberOfDayToBloom(bloomDay)
    maxPossibleDays := findMaximumNumberOfDayToBloom(bloomDay)

    for minPossibleDays <= maxPossibleDays {
        possibleDays := minPossibleDays + (maxPossibleDays - minPossibleDays) / 2
        bouquetsPossible := findNumberOfBouquestMade(bloomDay, possibleDays, k)

        if bouquetsPossible < m {
            minPossibleDays = possibleDays + 1
        } else {
            maxPossibleDays = possibleDays - 1
        }
    }

    return minPossibleDays
}

func findNumberOfBouquestMade(bloomDay []int, days int, k int) int {
    bouquets := 0
    flowers := 0

    for _, bloom := range bloomDay {
        if bloom <= days {
            flowers++
            if flowers == k {
                bouquets++
                flowers = 0
            }
        } else {
            flowers = 0
        }
    }

    return bouquets
}

func findMaximumNumberOfDayToBloom(bloomDay []int) int {
    max := math.MinInt32

    for i := range bloomDay {
        if max < bloomDay[i] {
            max = bloomDay[i]
        }
    }

    return max
}

func findMinimumNumberOfDayToBloom(bloomDay []int) int {
      min := math.MaxInt32

    for i := range bloomDay {
        if min > bloomDay[i] {
            min = bloomDay[i]
        }
    }

    return min
}

