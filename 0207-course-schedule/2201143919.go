func canFinish(numCourses int, prerequisites [][]int) bool {
    g := NewGraph(prerequisites)
    visited := make([]bool, numCourses)
    pathVisited := make([]bool, numCourses)

    for idx := range numCourses {
        if !visited[idx] {
            flag := dfs(g, idx, visited, pathVisited)
            if !flag {
                return flag
            }
        }
      
    }
    
    return true
}

type Graph struct {
    adj map[int]map[int]struct{}
}

func NewGraph(pr [][]int) *Graph {
    adj := make(map[int]map[int]struct{})

    for _ ,element := range pr {
        first := element[0]
        second := element[1]

        if _, ok := adj[second]; !ok {
            adj[second] = make(map[int]struct{})
        }

        adj[second][first] = struct{}{}
    }

    return &Graph{
        adj: adj,
    }
}

func dfs(g *Graph, idx int, visited []bool, pathVisited []bool) bool { 
    visited[idx] = true
    pathVisited[idx] = true

    for key := range g.adj[idx] {
        if !visited[key] {
            flag := dfs(g, key, visited, pathVisited)
            if !flag {
                return flag
            }    
        } else if pathVisited[key] {
            return false
        } 
    }

    pathVisited[idx] = false
    return true
}


