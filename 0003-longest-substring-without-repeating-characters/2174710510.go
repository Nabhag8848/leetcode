func lengthOfLongestSubstring(s string) int {
    hash_map := make(map[byte]int)
    
    l := 0
    r := 0

    maxi := 0 

    for r < len(s) {
        _, ok := hash_map[s[r]]

        if !ok {
            hash_map[s[r]] = r
            maxi = max(maxi, r - l + 1)
            r++
        } else  {
            left, _ := hash_map[s[l]]
            delete(hash_map, s[l])
            l = left + 1
        }
    }

    return maxi
}