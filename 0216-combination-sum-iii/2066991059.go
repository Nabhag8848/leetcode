func combinationSum3(k int, n int) [][]int {
    return helper(k, n, 1, 0, []int{})
}

func helper(length int, sum_upto int, num int, sum int, arr []int) [][]int{

    if sum == sum_upto && length == len(arr){
        c := make([]int, len(arr))
        copy(c, arr)
        return [][]int{c}
    }

    var result [][]int

    for i:=num; i <= 9; i++ {
        if len(arr) > length || sum > sum_upto {
            break
        }  
        
        result = append(result, helper(length, sum_upto, i + 1, sum + i, append(arr, i))...)
    }

    return result
}