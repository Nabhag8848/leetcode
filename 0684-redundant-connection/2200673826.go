func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    g := NewGraph()

    for _, edge := range edges {
        u, v := edge[0], edge[1]

        visited := make([]bool, n+1)

        if dfs(g, visited, u, v) {
            return edge
        }

        g.addEdge(u, v)
    }

    return []int{}
}

type Graph struct {
    adj map[int]map[int]struct{}
}

func NewGraph() *Graph {
    return &Graph{adj: make(map[int]map[int]struct{})}
}

func (g *Graph) addEdge(u, v int) {
    if g.adj[u] == nil {
        g.adj[u] = make(map[int]struct{})
    }
    if g.adj[v] == nil {
        g.adj[v] = make(map[int]struct{})
    }

    g.adj[u][v] = struct{}{}
    g.adj[v][u] = struct{}{}
}

func dfs(g *Graph, visited []bool, idx, target int) bool {
    if idx == target {
        return true
    }

    visited[idx] = true

    for node := range g.adj[idx] {
        if !visited[node] && dfs(g, visited, node, target) {
            return true
        }
    }

    return false
}
