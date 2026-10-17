func lengthOfLongestSubstring(s string) int {
    hash_map := make(map[byte]struct{})
    
    l := 0
    r := 0

    maxi := 0 

    for r < len(s) {
        _, ok := hash_map[s[r]]

        if !ok {
            hash_map[s[r]] = struct{}{}
            r++
            maxi = max(maxi, r - l)
        } else  {
            delete(hash_map, s[l])
            l++
        }
    }

    return maxi
}