func findDegrees(matrix [][]int) []int {
    degrees := make([]int, len(matrix))

    for i, row := range matrix {
        for j := 0; j < i; j++ {
            if row[j] == 1 {
                degrees[i]++
                degrees[j]++
            }
        }
    }

    return degrees
}
