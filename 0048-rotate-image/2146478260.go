func rotate(matrix [][]int)  {
    // transpose

    for i := range matrix {
        for j := range matrix[i] {
            if i > j {
                tmp := matrix[i][j]
                matrix[i][j] = matrix[j][i]
                matrix[j][i] = tmp
            }
            
        }
    }

    for row,_ := range matrix {
        l := 0
        r := len(matrix[row]) - 1

        for l < r {
            tmp := matrix[row][l]
            matrix[row][l] = matrix[row][r]
            matrix[row][r] = tmp
            l++
            r--
        }
    }
}

/*  
    i,j

    1 4 7
    2 5 8
    3 6 9

    i,j -> j,i

    5 2 13 15
    1 4 6 1
*/