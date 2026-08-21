func firstUniqChar(s string) int {
    var freq [26]int

    for i := range s {
        freq[s[i]-'a']++
    }

    for i := range s {
        if freq[s[i] - 'a'] == 1 {
            return i
        }
    }

    return -1
}