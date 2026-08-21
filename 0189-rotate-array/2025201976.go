func rotate(nums []int, k int)  {
    k_elements := k % len(nums) 
    nums = reverse(nums, 0, len(nums) - 1)
    nums = reverse(nums, 0, k_elements - 1)
    nums = reverse(nums, k_elements, len(nums) - 1) 
}

func reverse(nums []int, start int, end int) []int {
    i, j := start, end
    for i < j {
        x:= nums[i]
        nums[i] = nums[j]
        nums[j] = x

        i++
        j-- 
    }

    return nums
}