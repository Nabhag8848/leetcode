func largestOddNumber(num string) string {
    for i := len(num) - 1; i >= 0;i-- {
        if isOdd(num[i]) {
            return num[:i + 1]
        }
    }

    return ""
}

func isOdd(num byte) bool {
    switch num {
        case '1','3','5','7','9':
            return true
        default:
            return false
    }

}