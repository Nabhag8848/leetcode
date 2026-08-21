func combinationSum(candidates []int, target int) [][]int {
    return helper(candidates, 0, target, []int{})
}

func helper(candidates []int, index int, target int, arr []int) [][]int {

    if index == len(candidates) {
        if target == 0 {
            c := make([]int, len(arr))
            copy(c, arr)
            return [][]int{c}
        }

        return [][]int{}
    }
    
    var result [][]int

    if target > 0 {
        result = append(result, helper(candidates, index, target - candidates[index], append(arr, candidates[index]))...)
    }

    result = append(result, helper(candidates, index + 1, target, arr)...)

    return result
}