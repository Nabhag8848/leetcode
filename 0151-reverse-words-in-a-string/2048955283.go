func reverseWords(s string) string {
    s = strings.TrimSpace(s)

    str_arr := make([]string, 0)

    var builder strings.Builder

    for i := range s {
        if s[i] == ' ' {
            str_arr = append(str_arr, builder.String())
            builder.Reset()
            for i < len(s) && s[i] == ' ' {
                i++
            }
        } else {
            builder.WriteByte(s[i])
        }
    }

    str_arr = append(str_arr, builder.String())
    builder.Reset()

    for i := len(str_arr) - 1; i >= 0;i-- {
        builder.WriteString(str_arr[i])
        if i != 0 && len(str_arr[i]) > 0 {
            builder.WriteString(" ")
        }
    }

    return builder.String()

}