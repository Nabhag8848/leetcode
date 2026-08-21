func longestCommonPrefix(strs []string) string {
    var builder strings.Builder
    j := 0

    for {
        if j >= len(strs[0]) {
            break
        }
        char := string(strs[0][j])

        isBreak := false
        for _, s := range strs[1:] {
            if j >= len(s) || string(s[j]) != char {
                isBreak = true
                break
            }
        }

        if isBreak {
            break
        }

        builder.WriteString(char)
        j++
    }

    return builder.String()
}