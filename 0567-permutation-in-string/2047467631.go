func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }

    var freq1, freq2 [26]int

    for i := 0; i < len(s1); i++ {
        freq1[s1[i]-'a']++
        freq2[s2[i]-'a']++
    }

    if freq1 == freq2 {
        return true
    }

    for i := 1; i <= len(s2)-len(s1); i++ {
        freq2[s2[i-1]-'a']--
        freq2[s2[i+len(s1)-1]-'a']++
        if freq1 == freq2 {
            return true
        }
    }

    return false
}