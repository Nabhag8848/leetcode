func jump(nums []int) int {
    result := 0 
    l := 0
    r := 0


    for r < len(nums) - 1 {
        far := 0
        
        for i := l; i <= r; i++ {
            far = max(far, i + nums[i])
        }

        l = r + 1
        r = far
        result++
    }

    return result
}
