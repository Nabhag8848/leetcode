func partition(s string) [][]string {
    return helper(s, 0, []string{})
}

func helper(s string, index int, res []string) [][]string {
    if index == len(s) {
        c := make([]string, len(res))
        copy(c, res)
        return [][]string{c}
    }

    var result [][] string
    for i:=index; i < len(s); i++ {
        if ok := palindrome(s, index, i); ok {
            result = append(result, helper(s, i + 1, append(res, s[index:i + 1]))...)
        }
    }

    return result
}

func palindrome(s string, start int, end int) bool {
    for start <= end {
        if s[start] != s[end] {
            return false
        }
        start++
        end--
    }

    return true
}