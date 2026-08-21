func maxSubArray(nums []int) int {
    maxEnding, result := nums[0], nums[0]
    
    for i:=1; i <len(nums);i++ {
        maxEnding = int(math.Max(float64(maxEnding + nums[i]), float64(nums[i])))
        result = int(math.Max(float64(result), float64(maxEnding)))
    }

    return result
}