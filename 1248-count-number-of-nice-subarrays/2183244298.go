func numberOfSubarrays(nums []int, k int) int {
    return niceSubarrayLessEqual(nums, k) - niceSubarrayLessEqual(nums, k - 1)
}

func niceSubarrayLessEqual(nums []int, k int) int {
    result := 0

    l := 0
    r := 0

    count := 0

    for r < len(nums) {
        is_right_odd := (nums[r] & 1) == 1

        if is_right_odd {
            count++
        }

        for count > k {
            is_left_odd := (nums[l] & 1) == 1 

            if is_left_odd {
                count--
            }

            l++ 
        }

        result = result + (r - l + 1)
        r++
    }

    return result
}
