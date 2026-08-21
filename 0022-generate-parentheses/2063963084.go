func generateParenthesis(n int) []string {
    return helper("", 0, 0, n)
}

func helper(current string, open int, close int, n int) []string {
    if len(current) == 2*n {
        return []string{current}
    }

    var result []string

    if open < n {
        result = append(result, helper(current+"(", open+1, close, n)...)
    }
    if close < open {
        result = append(result, helper(current+")", open, close+1, n)...)
    }

    return result
}