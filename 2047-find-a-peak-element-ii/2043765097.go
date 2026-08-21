func findPeakGrid(mat [][]int) []int {
   low := 0 
   high := len(mat[0]) - 1

   for low <= high {
      mid := low + (high - low) / 2
      row := findMaxElementRow(mat, mid)
      left := math.MinInt32
      right := math.MinInt32
      
      if mid > 0 {
        left = mat[row][mid - 1]
      }

      if mid < len(mat[0]) - 1 {
        right = mat[row][mid + 1]
      }

      if mat[row][mid] > left && mat[row][mid] > right {
          return []int{row, mid}
      } else if mat[row][mid] > right {
        high = mid - 1
      } else {
        low = mid + 1
      }
   }

   return []int{-1,-1}
}

func findMaxElementRow(mat [][]int, col int) int {
    max := math.MinInt32
    maxIndex := -1

    for i := range mat {
        if max < mat[i][col] {
            max = mat[i][col]
            maxIndex = i
        }
    }

    return maxIndex
}
