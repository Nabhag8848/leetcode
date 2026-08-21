func findPeakGrid(mat [][]int) []int {
    hash_map := make(map[int]int)
    row_col_map := make(map[int]int)
    for row := 0; row < len(mat); row++ {
        col_index := peak_element(mat, row)
        if col_index != - 1 {
            value, _ := hash_map[col_index]
            if value < mat[row][col_index] { 
                hash_map[col_index] = mat[row][col_index]
                row_col_map[col_index] = row
            }
        }
    }

   for col := range hash_map {
    row := row_col_map[col]

    val := mat[row][col]

    ok := true

    if row > 0 && mat[row-1][col] >= val {
        ok = false
    }
    if row+1 < len(mat) && mat[row+1][col] >= val {
        ok = false
    }
    if col > 0 && mat[row][col-1] >= val {
        ok = false
    }
    if col+1 < len(mat[0]) && mat[row][col+1] >= val {
        ok = false
    }

    if ok {
        return []int{row, col}
    }
}

    return []int{-1, -1}
}

func peak_element(mat [][]int, row int) int {
    low := 0
    high := len(mat[0]) - 1

    if low == high {
        return low
    }   

    if mat[row][low] > mat[row][low + 1] {
        return low
    }
    
    if mat[row][high] > mat[row][high - 1] {
        return high
    }

    low++
    high--

    for low <= high {
        mid := low + (high - low) / 2 

        if mat[row][mid] > mat[row][mid - 1] && mat[row][mid] > mat[row][mid + 1] { 
            return mid
        } else if mat[row][mid] < mat[row][mid - 1] { 
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return -1
}