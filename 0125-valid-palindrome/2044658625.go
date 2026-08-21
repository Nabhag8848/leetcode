func isPalindrome(s string) bool {
    s = strings.TrimSpace(s)
    var builder strings.Builder

    for i := range s {
        if (s[i] >= 65 && s[i] <= 90) {
            builder.WriteString(string(s[i] + 32))
        } else if (s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57) {
            builder.WriteString(string(s[i]))
        } 
    }

    str := builder.String()
    bytes := []byte(str)

    i := 0 
    j := len(bytes) - 1

    for i < j {
        if bytes[i] != bytes[j] {
            return false
        }

        i++
        j--
    }

    return true
}