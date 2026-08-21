func duplicateNumbersXOR(nums []int) int {
    seen := make(map[int]bool)
    result := 0
    for _, n := range nums {
        if seen[n] {
            result ^= n
        } else {
            seen[n] = true
        }
    }
    return result
}