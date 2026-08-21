func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    i := 0
    is_one_exist := false

    for i < len(nums) {
        for i < len(nums) && nums[i] == 0 {
            i++
        }

        j := i
        count := 0
        for j < len(nums) && nums[j] == 1 {
            j++
            count++
            is_one_exist = true
        }

        i = j
        if (count > max) {
            max = count
        }
    }

    if (is_one_exist) {
        return max
    }
    
    return 0
}