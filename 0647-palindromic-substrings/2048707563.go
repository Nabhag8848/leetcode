func countSubstrings(s string) int {
  count := 0

    for i:=0;i < len(s);i++ {
        checkPalindromes(s, i, i, &count)
        checkPalindromes(s, i, i + 1, &count)
    }

    return count
}

func checkPalindromes(s string, left int, right int, count *int) {

    for left >=0 && right < len(s) && s[left] == s[right] {
        *count = *count + 1
        left--
        right++
    } 
}