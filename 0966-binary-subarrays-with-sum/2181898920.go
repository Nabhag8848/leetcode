func numSubarraysWithSum(nums []int, goal int) int {
    return sumLessEqualGoal(nums, goal) - sumLessEqualGoal(nums, goal - 1)
}

func sumLessEqualGoal(nums []int , goal int) int {
    if goal < 0 {
        return 0
    }


    l := 0
    r := 0

    sum := 0
    result := 0

    for r < len(nums) {
        sum += nums[r]

        for sum > goal {
            sum -= nums[l]
            l++
        }

        result = result + (r - l + 1)
        r++
    }

    return result
}


  