func smallestDivisor(nums []int, threshold int) int {
    minPossDiv := 1
    maxPossDiv := findMaximumNumber(nums)

    for minPossDiv <= maxPossDiv {
        possibleDiv := minPossDiv + (maxPossDiv - minPossDiv) / 2
        sumOfDivisorResult := findSumOfDivisorResult(nums, possibleDiv)

        if sumOfDivisorResult <= threshold {
            maxPossDiv = possibleDiv - 1
        } else {
            minPossDiv = possibleDiv + 1
        }
    }

    return minPossDiv
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

func findSumOfDivisorResult(nums []int, divisor int) int {
    sum := 0

    for i := range nums {
        sum += int(math.Ceil(float64(nums[i]) / float64(divisor)))
    }

    return sum
}