func longestCommonPrefix(strs []string) string {
    prefix := strs[0]

    for _, s := range strs[1:] {
        for len(prefix) > 0 && (len(s) < len(prefix) || s[:len(prefix)] != prefix) {
            prefix = prefix[:len(prefix)-1]
        }
    }

    return prefix
}