func findCenter(edges [][]int) int {
    g := NewGraph(edges)
    n := len(g.adj)
    result := 1

    for node, edges := range g.adj {
        if len(edges) == n - 1 {
            result = node
            break
        }
    }

    return result
}

type Graph struct {
    adj map[int]map[int]struct{}
}

func NewGraph(edges[][]int) *Graph {
    adj := make(map[int]map[int]struct{})

    for _, edge := range edges {
        _, left_ok := adj[edge[0]]
        _, right_ok := adj[edge[1]]

        if !left_ok {
            adj[edge[0]] = make(map[int]struct{})
        } 

        if !right_ok {
            adj[edge[1]] = make(map[int]struct{})
        }

        adj[edge[0]][edge[1]] = struct{}{}
        adj[edge[1]][edge[0]] = struct{}{}
    }

    return &Graph{
        adj: adj,
    }
}
