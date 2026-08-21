func lengthOfLongestSubstring(s string) int {
    maxi := 0
    for i:=0; i < len(s);i++ {
        hash_arr := make([]int, 256)
        for j:=i; j < len(s); j++ {
            if hash_arr[s[j]] == 1 {
                break
            }

            maxi = max(j - i + 1, maxi)
            hash_arr[s[j]]++          
        }
    }

    return maxi
}