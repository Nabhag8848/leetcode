func shipWithinDays(weights []int, days int) int {
    minimumCapacity := findMaximumNumber(weights)
    maximumCapacity := sumOfWeights(weights)

    for minimumCapacity <= maximumCapacity {
        possibleCapacity := minimumCapacity + (maximumCapacity - minimumCapacity) / 2
        daysRequired := findShippingDaysWithCapacity(weights, possibleCapacity)

        if daysRequired > days {
            minimumCapacity = possibleCapacity + 1
        } else {
            maximumCapacity = possibleCapacity - 1
        }
    }

    return minimumCapacity
}

func findShippingDaysWithCapacity(weights []int, possibleCapacity int) int {
    totalDays := 1
    leftOverCapacity := possibleCapacity

    for i := range weights {
        if leftOverCapacity - weights[i] >= 0 {
            leftOverCapacity = leftOverCapacity - weights[i]
        } else {
            leftOverCapacity = possibleCapacity - weights[i]
            totalDays++
        }
    }   

    return totalDays
}

func sumOfWeights(weights []int) int {
    sum := 0

    for i := range weights {
        sum += weights[i]
    }

    return sum
}

func findMaximumNumber(nums []int) int {
    max := math.MinInt32

    for i := range nums {
        if max < nums[i] {
            max = nums[i]
        }
    }

    return max
}

/*
    1,2,3,4,5,6,7,8,9,10

    i:= 10 -> 55

    10 -> 31
    10 -> 20 (days 4)
    10 -> 15 
    
    25 / 2


    1,2,3,1,1

    3 -> 8
     
    3 -> 4
    1 -> 4

    1 -> 2



*/

