func searchMatrix(matrix [][]int, target int) bool {
    row_len := len(matrix)
    col_len := len(matrix[0])
    low := 0
    high := row_len - 1
    target_row := -1

    for low <= high {
       mid := low + (high - low) / 2

       if target > matrix[mid][0] {
            if target > matrix[mid][col_len - 1] {
               low = mid + 1
            } else {
                target_row = mid 
                break
            }
       } else if target < matrix[mid][0] {
           high = mid - 1
       } else {
           target_row = mid
           break
       }
    }

    if target_row == -1 {
        return false
    }

    return binary_search(matrix, target_row, target)
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
