func validStrings(n int) []string {
    return helper("", n)
}

func helper(res string, n int) []string {
    if n == 0 {
        return []string{res}
    }

    res_len := len(res)
    var result []string

    if res_len == 0 || (res_len > 0 && res[res_len - 1] == '1') {
        result = append(result, helper(res + "0", n - 1)...)
    }

    result = append(result, helper(res + "1", n - 1)...)

    
    return result
}
