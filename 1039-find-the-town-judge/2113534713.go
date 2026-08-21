func findJudge(n int, trust [][]int) int {
    indegree := make([]int, n+1)
    outdegree := make([]int, n+1)

    for _, t := range trust {
        outdegree[t[0]]++
        indegree[t[1]]++
    }

    for person := 1; person <= n; person++ {
        if indegree[person] == n-1 && outdegree[person] == 0 {
            return person
        }
    }
    return -1
}