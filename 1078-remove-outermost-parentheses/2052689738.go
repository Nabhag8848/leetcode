func removeOuterParentheses(s string) string {
    startIndex := -1
    var builder strings.Builder
    count := 0
    for i := range s {
        if s[i] == '(' {
            if startIndex == -1 {
                startIndex = i
            }

            count++
        } else if s[i] == ')' {
            count--
        }

        if count == 0 {
            builder.WriteString(s[startIndex + 1:i])
            startIndex = -1
        }
    }

    return builder.String()
}