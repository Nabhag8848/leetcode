func rearrangeArray(nums []int) []int {
    arr := make([]int, len(nums))
    j,k:=0, 1
    for i := 0; i < len(nums); i++ {
        if nums[i] < 0 {
            arr[k] = nums[i]
            k = k + 2
        } else {
            arr[j] = nums[i]
            j = j + 2

        }
    }

    return arr
}