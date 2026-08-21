func rotate(matrix [][]int)  {

    for i:=0; i < len(matrix); i++ {
        for j := 0; j < len(matrix) / 2; j++ {
            x := matrix[j][i]
            matrix[j][i] = matrix[(len(matrix) - 1) - j][i]
            matrix[(len(matrix) - 1) - j][i] = x
        }
    }

     for i:=0; i < len(matrix); i++ {
        for j := 0; j < len(matrix); j++ {

            if i < j {
                x := matrix[i][j]
                matrix[i][j] = matrix[j][i]
                matrix[j][i] = x
            }

        }
    }

        

  /*

        7 8 9
        4 5 6
        1 2 3

    
    1  2  3  4 
    5  6  7  8 
    9  10 11 12
    13 14 15 16

    13 9 5 1
    14 10 6 2
    15 11 7 3
    16 12 8 4

    00 -> 30 
    10 -> 31
    20 -> 32
    30 -> 33

    01 -> 20
    11 -> 21
    21 -> 22
    31 -> 23

    arr[i][n-j] = arr[j][i] 
    */
}