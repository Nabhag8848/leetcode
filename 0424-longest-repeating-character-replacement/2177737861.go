func characterReplacement(s string, k int) int {
    hash_map := make(map[byte]int)
    l := 0
    r := 0
    max_freq := 0
    result := 0

    for r < len(s) {

        hash_map[s[r]]++
        max_freq = max(max_freq, hash_map[s[r]])

        can_flip := (r - l + 1) - max_freq

        if can_flip > k {
            hash_map[s[l]]--
            if val, _ := hash_map[s[l]]; val == 0 {
                delete(hash_map, s[l])
            }
            l++
        } else {
            result = max(result, r - l + 1)
        }

        r++
    }

    return result
}