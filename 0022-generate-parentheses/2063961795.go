func generateParenthesis(n int) []string {
    var ans []string
    helper("", 0, 0, n, &ans)
    return ans
}

func helper(current string, open int, close int, n int, ans *[]string) {
    if len(current) == 2*n {
        *ans = append(*ans, current)
        return
    }
    if open < n {
        helper(current+"(", open+1, close, n, ans)
    }
    if close < open { 
        helper(current+")", open, close+1, n, ans)
    }
}