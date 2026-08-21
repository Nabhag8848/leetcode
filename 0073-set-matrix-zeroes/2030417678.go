func setZeroes(matrix [][]int)  {
    row := make([]int, len(matrix))
    col := make([]int, len(matrix[0]))

    for i:= 0; i < len(matrix); i++ {
        for j:=0; j < len(matrix[i]); j++{
            if matrix[i][j] == 0 {
                row[i] = -1
                col[j] = -1
            }
        }
    }


    for i:= 0; i < len(matrix); i++ {
        for j:=0; j < len(matrix[i]); j++{
            if row[i] == -1 ||  -1 == col[j] {
                matrix[i][j] = 0
            }
        }
    }

}
