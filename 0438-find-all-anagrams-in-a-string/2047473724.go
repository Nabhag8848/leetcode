func findAnagrams(s string, p string) []int {
     if len(s) < len(p) {
        return []int{}
    }

    var freq1, freq2 [26]int
    result := make([]int, 0)

    for i := 0; i < len(p); i++ {
        freq1[p[i]-'a']++
        freq2[s[i]-'a']++
    }

    if freq1 == freq2 {
        result = append(result, 0)
    }

    for i := 1; i <= len(s)-len(p); i++ {
        freq2[s[i-1]-'a']--
        freq2[s[i+len(p)-1]-'a']++
        if freq1 == freq2 {
            result = append(result, i)
        }
    }

    return result
}