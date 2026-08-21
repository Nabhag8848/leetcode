func numJewelsInStones(jewels string, stones string) int {
    hash_map := make(map[int32]struct{})
    count := 0

    for _, value := range jewels {
        hash_map[value] = struct{}{}
    }

    for _, value := range stones {
        if _, is_exist := hash_map[value]; is_exist {
            count++
        }
    }

    return count
}