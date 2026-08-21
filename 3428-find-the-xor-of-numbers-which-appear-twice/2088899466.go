func duplicateNumbersXOR(nums []int) int {
    seen := 0
    xor := 0

    for i := range nums {
        bit := 1 << nums[i]
        if (seen & bit) != 0 {
            xor ^= nums[i]
        } else {
            seen = seen ^ (1 << nums[i])
        }
    }

    return xor
}