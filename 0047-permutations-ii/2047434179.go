func permuteUnique(nums []int) [][]int {
    ans := make([][]int, 0)
    
    recursive(0, nums, &ans)
    return ans
}

func recursive(idx int, nums []int, ans *[][]int) {
     if idx == len(nums) {
        temp := make([]int, len(nums))
        copy(temp, nums)
        *ans = append(*ans, temp)
        return 
     }

    seen := map[int]bool{}


     for i := idx; i < len(nums); i++ {
         if seen[nums[i]] {
            continue
        }
        seen[nums[i]] = true
        nums[idx], nums[i] = nums[i], nums[idx]
        recursive(idx + 1, nums, ans)
        nums[idx], nums[i] = nums[i], nums[idx]
     }
}


