func subsetsWithDup(nums []int) [][]int {
    memo := make(map[string]struct{})
    return helper(nums, 0, memo, []int{})
}

func helper(nums []int, index int, memo map[string]struct{}, arr []int) [][]int{
    if index == len(nums) {
        c := make([]int, len(arr))
        copy(c, arr)
        sort.Ints(c)
        str := fmt.Sprintf("%v", c)

        if _,ok := memo[str]; !ok {
            memo[str] = struct{}{}
            return [][]int{c}
        }

        return [][]int{}
    }

    var result [][]int
    result = append(result, helper(nums, index + 1, memo, arr)...)
    result = append(result, helper(nums, index + 1, memo, append(arr, nums[index]))...)

    return result
}
