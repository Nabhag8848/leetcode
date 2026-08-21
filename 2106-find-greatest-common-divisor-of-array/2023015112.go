func findMinMax(nums []int) (int, int) {
    min := math.MaxInt32
    max := math.MinInt32

    for x := range nums {
        if nums[x] < min {
            min = nums[x]
        }
    }

    for x := range nums {
        if nums[x] > max {
            max = nums[x]
        }
    }

    return min, max
}

func findGCD(nums []int) int {
    x, y := findMinMax(nums)

    for int(math.Min(float64(x), float64(y))) != 0 {
        if (x <= y) {
            y = y - x
        } else {
            x = x - y
        }
    }

    return int(math.Max(float64(x), float64(y)))
}