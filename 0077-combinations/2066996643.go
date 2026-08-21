func combine(n int, k int) [][]int {
    return helper(k, n, 1, []int{})
}

func helper(length int, upto int, num int, arr []int) [][]int{
    if length == len(arr){
        c := make([]int, len(arr))
        copy(c, arr)
        return [][]int{c}
    }

    var result [][]int

    for i:=num; i <= upto; i++ {
        if len(arr) > length {
            break
        }  
        
        result = append(result, helper(length, upto, i + 1, append(arr, i))...)
    }

    return result
}