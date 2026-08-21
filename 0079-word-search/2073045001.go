func exist(board [][]byte, word string) bool {
    m := len(board)
    n := len(board[0])
    visited := make([][]bool, m)

    for i := range m {
        visited[i] = make([]bool, n)
    }

    for i := range m {
        for j := range n {
            visited[i][j] = false
        }
    }

    for i := range m {
        for j := range n {
            if word[0] == board[i][j] {
                visited[i][j] = true
                if result := path(board, word, visited, 1, i, j, m - 1, n - 1); result {
                    return true
                } 

                visited[i][j] = false
            }

        }
    }

    return false

}

func path(board [][]byte, word string, visited [][]bool, index int, row int, col int, m int, n int) bool {
    if index == len(word) {
        return true
    }

    if row < m && board[row + 1][col] == word[index] && !visited[row + 1][col] {
        visited[row + 1][col] = true
        if ok := path(board, word, visited, index + 1, row + 1, col, m, n); ok {
            return ok
        }

        visited[row + 1][col] = false
    }

    if col < n && board[row][col + 1] == word[index] && !visited[row][col + 1]{
        visited[row][col + 1] = true
        if ok := path(board, word, visited, index + 1, row, col + 1, m, n); ok {
            return ok
        }
        visited[row][col + 1] = false
    }

    if row > 0 && board[row - 1][col] == word[index] && !visited[row - 1][col] {
        visited[row - 1][col] = true
        if ok := path(board, word, visited, index + 1, row - 1, col, m, n); ok {
            return ok
        }
        visited[row - 1][col] = false
    }

    if col > 0 && board[row][col - 1] == word[index] && !visited[row][col - 1]{
        visited[row][col - 1] = true
        if ok := path(board, word, visited, index + 1, row, col - 1, m, n); ok {
            return ok
        }
        visited[row][col - 1] = false
    }

    return false
}