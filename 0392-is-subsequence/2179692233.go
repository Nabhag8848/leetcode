func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }

    i := 0
    j := 0

    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i++
        }

        j++
    }

    if i == len(s) {
        return true
    }

    return false
}