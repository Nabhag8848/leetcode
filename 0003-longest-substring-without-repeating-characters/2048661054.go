func lengthOfLongestSubstring(s string) int {
    hash_arr := make([]int, 256)
    for i := range hash_arr {
        hash_arr[i] = -1
    }

    left := 0
    right := 0
    maxi := 0

    for right < len(s) {
        if hash_arr[s[right]] != -1 {
            if hash_arr[s[right]] >= left {
                left = hash_arr[s[right]] + 1
            }
        }

        maxi = max(right - left + 1, maxi)
        hash_arr[s[right]] = right
        right++
    }

    return maxi
}