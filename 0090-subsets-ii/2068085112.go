func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    return helper(nums, 0, []int{})
}

func helper(candidates []int, index int, arr []int) [][]int {

    if index == len(candidates) {
        res := make([]int, len(arr))
        copy(res, arr)
        return [][]int{res}
    }

    var result [][]int
    c := make([]int, len(arr))
    copy(c, arr)
    result = append(result, c)

    for i:=index; i < len(candidates); i++ {
        if i > index && candidates[i] == candidates[i - 1] {
            continue
        }
       
        result = append(result, helper(candidates, i + 1, append(arr, candidates[i]))...)
    }

    return result
}