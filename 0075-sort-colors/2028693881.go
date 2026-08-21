func sortColors(nums []int)  {
    counts := []int{0, 0, 0}

    for i:=0;i<len(nums);i++{
        counts[nums[i]]++
    }

    for i:=0 ;i<len(nums);i++{
        if i < counts[0] {
            nums[i] = 0
        }

        if i >= counts[0] && i < counts[1] + counts[0] {
            nums[i] = 1
        }

        if i >= counts[0] + counts[1] && i < counts[2] + counts[1] + counts[0] {
            nums[i] = 2
        }
    }
}