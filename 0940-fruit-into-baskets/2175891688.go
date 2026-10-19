func totalFruit(fruits []int) int {
    hash_map := make(map[int]int)

    result := 0
    l := 0
    r := 0 

    for r < len(fruits) {
        _, ok := hash_map[fruits[r]]

        if !ok {
            if len(hash_map) < 2 {
                hash_map[fruits[r]]++
                r++
            } else {
                hash_map[fruits[l]]--
                if val, _ := hash_map[fruits[l]]; val == 0 {
                    delete(hash_map, fruits[l])
                }
                l++
            }
        } else {
            hash_map[fruits[r]]++
            r++
        }

        result = max(result, r - l)
    }

    return result
}