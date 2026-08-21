func spiralOrder(matrix [][]int) []int {
    lRow, hRow := 0, len(matrix) - 1
    lCol, hCol := 0, len(matrix[0]) - 1
    result := make([]int, 0)

    for lRow <= hRow && lCol <= hCol {
        for i:= lCol; i <= hCol; i++ {
            result = append(result, matrix[lRow][i])
        }

        lRow++

        for i:=lRow; i <= hRow; i++ {
            result = append(result, matrix[i][hCol])
        }
        
        hCol--

        if lRow <= hRow {
            for i:=hCol; i >= lCol; i-- {
                result = append(result, matrix[hRow][i])
            }

            hRow--
        }

        if hCol >= lCol {
            for i:=hRow; i >= lRow; i-- {
               result = append(result, matrix[i][lCol])
        }   

            lCol++
        }
    }

    return result
}