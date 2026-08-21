func subsets(nums []int) [][]int {
    sort.Ints(nums)
    return helper(nums, 0, []int{})
}

func helper(nums []int, index int, res []int) [][]int {
    if index >= len(nums) {
        c := make([]int, len(res))
        copy(c, res)
        return [][]int{c}
    }

    result := make([][]int, 0)
    result = append(result, helper(nums, index + 1, res)...)
    result = append(result, helper(nums, index + 1, append(res, nums[index]))...)

    return result
}