func combinationSum4(candidates []int, target int) int {
    memo := make(map[int]int)
    return helper(candidates, target, memo)
}

func helper(candidates []int, target int, memo map[int]int) int {
    if target == 0 {
        return 1
    }

     if v, ok := memo[target]; ok {
        return v
    }

    var result int

    for i:=0; i < len(candidates); i++ {
        if target-candidates[i] >= 0 {
            result = result + helper(candidates, target - candidates[i], memo)
        }

    }

    memo[target] = result

    return result
}