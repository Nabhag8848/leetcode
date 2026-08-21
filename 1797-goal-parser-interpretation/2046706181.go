func interpret(command string) string {
    var builder strings.Builder
    i := 0

    for i < len(command) {
        if command[i] == 'G' {
            builder.WriteString("G")
            i = i + 1
        } else if command[i] == '(' {
            if command[i + 1] == ')' {
                builder.WriteString("o")
                i = i + 2
            } else {
                builder.WriteString("al")
                i = i + 4
            }
        }
    }

    return builder.String()
}