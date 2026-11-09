func checkValidString(s string) bool {
    n := len(s)

    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    return helper(s, 0, 0, dp)
}

func helper(s string, idx int, count int, dp [][]int) bool {
    if count < 0 {
        return false
    }

    if idx == len(s) {
        return count == 0
    }

    if dp[idx][count] != -1 {
        return dp[idx][count] == 1
    }

    var valid bool

    if s[idx] == '(' {
        valid = helper(s, idx+1, count+1, dp)
    } else if s[idx] == ')' {
        valid = helper(s, idx+1, count-1, dp)
    } else {
        valid = helper(s, idx+1, count-1, dp) ||
            helper(s, idx+1, count+1, dp) ||
            helper(s, idx+1, count, dp)
    }

    if valid {
        dp[idx][count] = 1
    } else {
        dp[idx][count] = 0
    }

    return valid
}
