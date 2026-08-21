func setZeroes(matrix [][]int)  {
    var isColZero = false

    for i:= 0; i < len(matrix); i++ {
       

        for j:=0; j < len(matrix[i]); j++{
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                if j == 0 {
                    isColZero = true
                }else {
                    matrix[0][j] = 0
                }
                
            }
        }
    }


    for i:= len(matrix) - 1; i > 0 ; i-- {
        for j:=len(matrix[i]) - 1; j > 0; j--{
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    if matrix[0][0] == 0 {
        for i:=0; i < len(matrix[0]); i++ {
            matrix[0][i] = 0
        }
    }

    if isColZero {
         for i:=0; i < len(matrix); i++ {
            matrix[i][0] = 0
        }
    }   
}
