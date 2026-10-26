func numberOfSubstrings(s string) int {
    hash_map := make(map[byte]int)

    l := 0
    r := 0
    result := 0

    for r < len(s) {
        hash_map[s[r]]++ 

        for len(hash_map) == 3 {
            result = result + (len(s) - r)

            hash_map[s[l]]--

            if val, _ := hash_map[s[l]]; val == 0 {
                delete(hash_map, s[l])
            }

            l++
        }

        r++
    }

    return result
}