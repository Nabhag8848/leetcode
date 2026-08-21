func countBits(n int) []int {
    result := make([]int, n+1)
    highestPowerOf2 := 1

    for i := 1; i <= n; i++ {
        if highestPowerOf2 * 2 == i {
            highestPowerOf2 = i
        }
        result[i] = 1 + result[i-highestPowerOf2]
    }

    return result
}