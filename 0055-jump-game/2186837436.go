func canJump(nums []int) bool {
    maxIdx := 0

    for i := range nums {
        if i > maxIdx {
            return false 
        }

        maxIdx = max(maxIdx, i + nums[i])
    }

    return true
}

