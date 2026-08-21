func gardenNoAdj(n int, paths [][]int) []int {
    edges := make(map[int]map[int]struct{})
    assignment := make(map[int]int)

    for i := range paths {
        x := paths[i][0]
        y := paths[i][1]

        if edges[x] == nil {
            edges[x] = make(map[int]struct{})
        }
        
        edges[x][y] = struct{}{}

        if edges[y] == nil {
            edges[y] = make(map[int]struct{})
        }
        edges[y][x] = struct{}{}
    }

    var result []int

    for i:=1; i <= n; i++ {
        
        for j := 1; j < 5; j++ {
            is_break := false
            for key, _ := range edges[i] {
                if assigned, ok := assignment[key]; ok {
                    if assigned == j {
                        is_break = true
                        break
                    }
                }
            }
            
            if !is_break {
                assignment[i] = j
                break
            }
        }
    }

    for i:=1; i <= n; i++ {
        result = append(result, assignment[i])
    }

    return result
}

/*
    1 2 3 

    1 - 2 - 3
      \ __ /


    1 - 2 3 4

       1 -> 1
    
*/