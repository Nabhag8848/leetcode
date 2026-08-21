func removeDuplicates(nums []int) int {
    if(len(nums) < 2) {
        return 1
    }

    count := 0
    i:=0
    for i < len(nums) {
        j := i+1;
        for  j < len(nums) && nums[i] == nums[j] {
            j++
        }

        count = count + 1
        if j < len(nums) {
            nums[count] = nums[j]
        }else {
            nums[count - 1] = nums[j - 1]
        }
        i = j
    }

    return count
}