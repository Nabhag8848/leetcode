func isAlienSorted(words []string, order string) bool {
    hash_map := make(map[byte]int)

    for i := range order {
        hash_map[order[i]] = i
    }

    for i := 0; i < len(words) - 1; i++ {
        str1 := words[i]
        str2 := words[i + 1]

        if str1 == str2 {
            continue
        }

        upto := min(len(str1), len(str2))
        idx := 0

        for idx < upto && str1[idx] == str2[idx]  {
            idx++
        }
        if idx >= upto {
            if len(str1) > len(str2) {
                return false
            }
        }

        if idx < upto && hash_map[str1[idx]] > hash_map[str2[idx]] {
            return false
        }
    }

    return true
}
