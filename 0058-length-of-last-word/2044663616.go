func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    count := 0

    for i:=len(s) - 1; i >= 0; i-- {
        if string(s[i]) == " " {
            break
        } else {
            count++
        }
    }

    return count
}