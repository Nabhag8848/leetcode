func check(nums []int) bool {
    var fault = false
    var len = len(nums)
    for i := 0; i < len; i++ {
        if(nums[i] > nums[(i + 1) % len]) {

            if(fault) {
                return false
            }
            fault = true
        }
    }

    return true
}
