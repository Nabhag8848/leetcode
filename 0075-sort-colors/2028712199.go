func sortColors(nums []int)  {
    dutchNationalFlagAlgorithm(nums)
}

func dutchNationalFlagAlgorithm(nums []int) {
    low := 0
    mid := 0
    high := len(nums) - 1

    for mid <= high {
        if nums[mid] == 0 {
            swap(nums, low, mid)
            low++
            mid++
        } else if nums[mid] == 1 {
            mid++
        } else { 
            swap(nums, high, mid)
            high--
        }
    }

}

func swap(nums []int, i, j int) {
    x := nums[i]
    nums[i] = nums[j]
    nums[j] = x
}

func firstAppraoch(nums []int) {
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