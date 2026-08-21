func longestPalindrome(s string) string {
    res := ""

    for i:=0;i < len(s);i++ {
        checkPalindromes(s, i, i, &res)
        checkPalindromes(s, i, i + 1, &res)
    }

    return res

}

func checkPalindromes(s string, left int, right int, res *string) {

    for left >=0 && right < len(s) && s[left] == s[right] {
        if right - left + 1 > len(*res) {
            *res = s[left:right+1]
        }
        left--
        right++
    } 
}