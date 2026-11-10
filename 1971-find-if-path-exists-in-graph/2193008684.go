func validPath(n int, edges [][]int, source int, destination int) bool {
    g := NewGraph(n, edges)
    visited := make([]bool, n)

    return helper(n, g, source, destination, &visited)
}

func helper(n int, g *Graph, source int, destination int, visited *[]bool) bool {
    if source == destination {
        return true
    }
    
    (*visited)[source] = true

    for node,_ := range g.adj[source] {
        if !(*visited)[node] {
            result := helper(n, g, node, destination, visited)
            if result {
                return true
            }
        }
    }

    return false
}

type Graph struct {
    adj map[int]map[int]struct{}
}

func NewGraph(n int, edges[][]int) *Graph {
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
