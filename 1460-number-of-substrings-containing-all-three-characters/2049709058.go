func numberOfSubstrings(s string) int {
    last_seen := []int{-1, -1, -1}
    sum := 0

    for i := range s {
        last_seen[s[i] - 'a'] = i

        if last_seen[0] != -1 && last_seen[1] != -1 && last_seen[2] != -1 {
            sum += min(last_seen[0], min(last_seen[1], last_seen[2])) + 1
        }
    }

    return sum
}

