func generate(numRows int) [][]int {
    result := make([][]int, 0, numRows)
    if numRows == 1 {
        return [][]int{{1}}
    }

    if numRows == 2 {
        return [][]int{{1},{1,1}}
    }

    result = append(result, []int{1})
    result = append(result, []int{1, 1})

    for i := 2; i < numRows; i++ {
        row := []int{}
        for j:=0; j <= i; j++ {
            if (j == 0 || j == i) {
                row = append(row, 1)
            } else {
                row = append(row, result[i - 1][j - 1] + result[i - 1][j])
            }   
        }

        result = append(result, row)
    }

    return result
}

 /*
    1
  1   1
1   2   1
1. 3  3    1

*/

// [1][0] + [1][1]
// [2][0] + [2][1]
// [2][1] + [2][2]

// [[1], [1,1], [1, 2, 1], [1, 1]]

// 1 + 2 .. + (n - 8)

