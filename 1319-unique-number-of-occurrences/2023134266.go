func uniqueOccurrences(arr []int) bool {
    hash_map := make(map[int]int)

    for x := range arr {
        hash_map[arr[x]] = hash_map[arr[x]] + 1
    }

    checker := make(map[int]bool)
    for _,v := range hash_map {
        _, ok := checker[v]

        if ok {
            return false
        }

        checker[v] = true
    }

    return true

}