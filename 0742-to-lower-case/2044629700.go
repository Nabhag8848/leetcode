func toLowerCase(s string) string {
    var builder strings.Builder

    for i := range s {
        if s[i] >= 65 && s[i] <= 90 {
            builder.WriteString(string(s[i] + 32))
        } else {
            builder.WriteString(string(s[i]))
        }
    }

    return builder.String()
}