func allPathsSourceTarget(graph [][]int) [][]int {
    var result [][]int
    dfs(graph, &result, []int{}, 0)
    return result
}

func dfs(graph [][]int, result *[][]int, source []int, curr int) {
    source = append(source, curr)

    if curr == len(graph) - 1 {
        arr := make([]int, len(source))
        copy(arr, source)
        *result = append(*result, arr)
    }

    for _, node := range graph[curr] {
        dfs(graph, result, source, node)    
    }

}
