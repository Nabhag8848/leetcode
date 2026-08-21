func rowAndMaximumOnes(mat [][]int) []int {
    number_of_ones := 0
    target_index := 0
    col_len := len(mat[0])

    for i := range mat {
	    sort.Ints(mat[i])
	}

    for i := range mat {
        if mat[i][0] == 1 && mat[i][col_len - 1] == 1 {
            number_of_ones = col_len
            target_index = i
            break
        } else if mat[i][0] == 0 && mat[i][col_len - 1] == 1 {
            count_ones := binary_search(mat, i)

            if count_ones > number_of_ones {
                number_of_ones = count_ones
                target_index = i
            }
        }
    }

    return []int{target_index, number_of_ones}
}

func binary_search(mat [][]int, row int) int {
    low := 0
    col_len := len(mat[row])
    high := col_len - 1

    for low <= high {
        mid := low + (high - low) / 2
        if mat[row][mid] == 1 {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return col_len - low
}

/*  
    0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1

*/