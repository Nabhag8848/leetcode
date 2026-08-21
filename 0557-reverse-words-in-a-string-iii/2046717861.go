func reverseWords(s string) string {
    var builder strings.Builder
    i := 0
    start_index := 0

    for i < len(s) {
        if s[i] == ' ' {
            target := reverse(s[start_index:i])
            builder.WriteString(target)
            builder.WriteString(" ")
            start_index = i + 1
        }

        i++
    }

    target := reverse(s[start_index:len(s)])
    builder.WriteString(target)

    return builder.String()
}

func reverse(s string) string {
    b := []byte(s)
    l, r := 0, len(b)-1
    for l < r {
        b[l], b[r] = b[r], b[l]
        l++
        r--
    }
    return string(b)
}