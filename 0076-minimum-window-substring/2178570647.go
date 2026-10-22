func minWindow(s string, t string) string {
    hash_map := make(map[byte]int)

    for i, _ := range t {
        hash_map[t[i]]++
    }

    l := 0
    r := 0

    count := 0
    min_len := math.MaxInt32
    s_idx := -1

    for r < len(s) {
        if val, _ := hash_map[s[r]]; val > 0 {
            count++
        }

        hash_map[s[r]]--

        for count == len(t) {
            if r - l + 1 < min_len {
                min_len = r - l + 1
                s_idx = l
            }

            hash_map[s[l]]++

            if hash_map[s[l]] > 0 {
                count = count - 1
            }

            l++
        }

        r++
    }

    if s_idx == -1 {
        return ""
    }
    end := s_idx + min_len 

    return s[s_idx:end]
}