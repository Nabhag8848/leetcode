func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    return helper(candidates, 0, target, []int{})
}

func helper(candidates []int, index int, target int, arr []int) [][]int {

    if target == 0 {
        res := make([]int, len(arr))
        copy(res, arr)
        return [][]int{res}
    }

    var result [][]int

    for i:=index; i < len(candidates); i++ {
        if i > index && candidates[i] == candidates[i - 1] {
            continue
        }

        if candidates[i] > target {
            break
        }

        result = append(result, helper(candidates, i + 1, target - candidates[i], append(arr, candidates[i]))...)
    }

    return result
}