func searchMatrix(matrix [][]int, target int) bool {
    col_len := len(matrix[0])
    for i := range matrix {
        if target >= matrix[i][0] {
            if target <= matrix[i][col_len - 1] {
                result := binary_search(matrix, i, target)
                if result {
                    return result
                }
            }
        } else if target < matrix[i][0] {
            break
        } 
    }

    return false
}

func binary_search(matrix [][]int, row int, target int) bool {
    low := 0
    col_len := len(matrix[row])
    high := col_len - 1

    for low <= high {
        mid := low + (high - low) / 2
        if target < matrix[row][mid] {
            high = mid - 1
        } else if target > matrix[row][mid] {
            low = mid + 1
        } else {
            return true
        }
    }

    return false
}