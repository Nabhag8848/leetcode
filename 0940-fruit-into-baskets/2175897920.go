func totalFruit(fruits []int) int {
    hash_map := make(map[int]int)

    result := 0
    l := 0
    r := 0 

    for r < len(fruits) {
        hash_map[fruits[r]]++

        if len(hash_map) > 2 {
            hash_map[fruits[l]]--
            if val, _ := hash_map[fruits[l]]; val == 0 {
                delete(hash_map, fruits[l])
            }
            l++
        }

        if len(hash_map) <= 2 {
            result = max(result, r - l + 1)
        }
        
        r++
    }

    return result
}