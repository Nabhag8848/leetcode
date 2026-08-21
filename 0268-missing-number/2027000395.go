func missingNumber(nums []int) int {
    var n  = len(nums)
    var sum = n * (n + 1) / 2
    var calculated_sum int

    for i := 0; i < len(nums); i++ {
        calculated_sum += nums[i]
    }

    return sum - calculated_sum
}