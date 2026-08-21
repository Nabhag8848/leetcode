func restoreString(s string, indices []int) string {
    char := []byte(s)

    for idx := range s {
        char[indices[idx]] = s[idx]
    }

    return string(char)
}