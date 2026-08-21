func moveZeroes(nums []int)  {
    i := 0
    j := 0

    for i < len(nums) && j < len(nums) {

        for i < len(nums) && nums[i] != 0 {
            i++
        }

        j = i + 1

        for j < len(nums) && nums[j] == 0 {
            j++
        }

        if (i < len(nums) && j < len(nums)) {
            nums[i] = nums[j]
            nums[j] = 0
        }

        i++
    }
    /* 
        0 0 0 1 2
        i 
              j
        
        1 2 0 0 0 
            i
                j

        1 2 3 4 5
                j

    */
}