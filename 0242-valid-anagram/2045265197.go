func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    var freq [26]int

    for i := range s {
        freq[s[i]-'a']++
    }

    for i := range t {
        freq[t[i]-'a']--
        if freq[t[i]-'a'] < 0 {
            return false
        }
    }

    return true
}