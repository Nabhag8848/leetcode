func findJudge(n int, trust [][]int) int {
    if len(trust) == 0 && n == 1 {
        return 1
    }

    hash_map := make(map[int]map[int]struct{})

    for i := range trust {
        if _, ok := hash_map[trust[i][0]]; !ok {
            hash_map[trust[i][0]] = make(map[int]struct{})
        }
        
        hash_map[trust[i][0]][trust[i][1]] = struct{}{}

    }

    possible := make(map[int]int, 0)

    for i := range trust {
        if _, ok := hash_map[trust[i][1]]; !ok {
            possible[trust[i][1]]++
        }
    }


    for k,v := range possible {
        if v == n - 1 {
            return k
        }
    }

    return -1

}