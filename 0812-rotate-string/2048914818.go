func rotateString(s string, goal string) bool {
    if len(s) < len(goal) {
        return false
    }

    s = s + s
    for i:=0; i <= len(s) - len(goal); i++ {
        if s[i:i + len(goal)] == goal {
            return true
        }
    }

    return false
}