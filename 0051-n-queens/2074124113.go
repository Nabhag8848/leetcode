func solveNQueens(n int) [][]string {
    visited := make([][]int, n)

    for i := range n {
        visited[i] = make([]int, n)
    }

    return helper(visited, 0, n, []string{})
}

func helper(visited [][]int, col, n int, res []string)[][]string {
    if len(res) == n {
        c := make([]string, n)
        copy(c, res)
        return [][]string{c}
    }

    var result [][]string

    for row := 0; row < n; row++ {
        if visited[row][col] == 0 {
            visited = markVisit(visited, row, col, n)
            var str string
            if row > 0 {
                str = strings.Repeat(".", row)
            }
            
            str += "Q"
            
            if n - row - 1 >= 0 {
                str += strings.Repeat(".", n - row - 1)
            } 

            result = append(result, helper(visited, col + 1, n , append(res, str))...)
            visited = undoVisit(visited, row, col, n)
        }
    }

    return result
}

func markVisit(visited [][]int, row, col, n int) [][]int {
    for i := range n {
        visited[i][col]++
        visited[row][i]++

        if row - i >= 0 && col - i >= 0 {
            visited[row - i][col - i]++
        }

        if row + i < n && col + i < n {
            visited[row + i][col + i]++
        }

        if row-i >= 0 && col+i < n {
            visited[row-i][col+i]++
        }
        if row+i < n && col-i >= 0 {
            visited[row+i][col-i]++
        }
    }

    return visited
}

func undoVisit(visited [][]int, row, col, n int) [][]int {
    for i := range n {
        visited[i][col]--
        visited[row][i]--

        if row - i >= 0 && col - i >= 0 {
            visited[row - i][col - i]--
        }

        if row + i < n && col + i < n {
            visited[row + i][col + i]--
        }

        if row-i >= 0 && col+i < n {
            visited[row-i][col+i]--
        }
        if row+i < n && col-i >= 0 {
            visited[row+i][col-i]--
        }
    }

    return visited
}