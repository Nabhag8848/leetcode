
func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    g := NewGraph(edges)

    for i := len(edges) - 1; i >= 0; i-- {
        u, v := edges[i][0], edges[i][1]

        delete(g.adj[u], v)
        delete(g.adj[v], u)

        visited := make([]bool, n+1)

        if dfs(g, visited, u, v) {
            return edges[i]
        }

        g.adj[u][v] = struct{}{}
        g.adj[v][u] = struct{}{}
    }

    return []int{}
}

type Graph struct {
    adj map[int]map[int]struct{}
}

func NewGraph(pr [][]int) *Graph {
    adj := make(map[int]map[int]struct{})

    for _, element := range pr {
        first := element[0]
        second := element[1]

        if _, ok := adj[first]; !ok {
            adj[first] = make(map[int]struct{})
        }

        adj[first][second] = struct{}{}

        if _, ok := adj[second]; !ok {
            adj[second] = make(map[int]struct{})
        }

        adj[second][first] = struct{}{}
    }

    return &Graph{adj: adj}
}

func dfs(g *Graph, visited []bool, idx, target int) bool {
    if idx == target {
        return true
    }

    visited[idx] = true

    for node := range g.adj[idx] {
        if !visited[node] {
            if dfs(g, visited, node, target) {
                return true
            }
        }
    }

    return false
}
