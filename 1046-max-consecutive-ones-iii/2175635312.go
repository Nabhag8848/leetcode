func longestOnes(nums []int, k int) int {
    left := 0
    maxLen := 0
    zeros := 0

    for right := range nums {
        if nums[right] == 0 {
            zeros++
        }

        for zeros > k {
            if nums[left] == 0 {
                zeros--
            }
            left++
        }

        maxLen = max(maxLen, right-left+1)
    }

    return maxLen
}