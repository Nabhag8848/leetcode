func countBits(n int) []int {
    var result []int
    for i := range n + 1 {
        result = append(result, helper(i))
    }

    return result
}

func helper(num int) int {
    count := 0
    for j := range 31 {
        bit := (1 << j) & num
        if bit != 0 {
            count++
        }
    }

    return count
}