func solveNQueens(n int) [][]string {
    cols := make([]bool, n)
    u_diagonal := make([]bool, 2 * n - 1)
    l_diagonal := make([]bool, 2 * n - 1)

    return helper(cols, u_diagonal, l_diagonal, 0, n, []string{})
}

func helper(cols, u_diagonal, l_diagonal []bool, col, n int, res []string)[][]string {
    if len(res) == n {
        c := make([]string, n)
        copy(c, res)
        return [][]string{c}
    }

    var result [][]string

    for row := 0; row < n; row++ {
        d1 := row + col
        d2 := row - col + n - 1

        if cols[row] || u_diagonal[d1] || l_diagonal[d2] {
            continue
        }

        cols[row], u_diagonal[d1], l_diagonal[d2] = true, true, true

        var str string
        if row > 0 {
            str = strings.Repeat(".", row)
        }
        str += "Q"
        if n-row-1 > 0 {
            str += strings.Repeat(".", n-row-1)
        }

        result = append(result, helper(cols, u_diagonal, l_diagonal, col+1, n, append(res, str))...)

        cols[row], u_diagonal[d1], l_diagonal[d2] = false, false, false
    }


    return result
}