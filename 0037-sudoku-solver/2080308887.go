func solveSudoku(board [][]byte)  {
    solve(board)
}

func solve(board [][]byte) bool {
    for i := range board {
        for j := range board[i] {

            if board[i][j] == '.' {
                for v := '1'; v <= '9'; v++ {
                    if ok := isPossible(board, i, j, byte(v)); ok {
                        board[i][j] = byte(v)

                        if isValid := solve(board); isValid {
                            return isValid
                        } else {
                            board[i][j] = '.'
                        }

                    }
                }

                return false
            }
        }
    }

    return true
}

func isPossible(board [][]byte, row, col int, char byte) bool {
    for i := 0; i < 9; i++ {
        if board[i][col] == char {
            return false
        }

        if board[row][i] == char {
            return false
        }

        boardRow := 3 * (row / 3) + i / 3
        boardCol := 3 * (col / 3) + i % 3 

        if board[boardRow][boardCol] == char {
            return false
        }
    }

    return true
}