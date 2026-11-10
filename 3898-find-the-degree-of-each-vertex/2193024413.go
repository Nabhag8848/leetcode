func findDegrees(matrix [][]int) []int {
    result := make([]int, len(matrix))

    for i := range matrix {
        for j := range matrix[i] {
            if i >= j && matrix[i][j] == 1 {
                result[i]++
                result[j]++
            }
        }    
    }

    return result
}
