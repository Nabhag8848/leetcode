func canJump(nums []int) bool {
    maxIdx := 0
    goal := len(nums) - 1

    for i := range nums {
        if i > maxIdx {
            return false 
        }

        maxIdx = max(maxIdx, i + nums[i])

        if maxIdx >= goal {
            return true
        }
    }

    return true
}

