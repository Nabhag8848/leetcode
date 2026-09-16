/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Level struct  {
    row int
    col int
}

type Nodes struct {
    data []int
}

func verticalTraversal(root *TreeNode) [][]int {
    hash_map := make(map[Level]Nodes)
    max_level := helper(root, 0, 0, hash_map) - 1 // zero index
    result := make([][]int, 0)

    fmt.Println(hash_map, max_level)

    for j := -max_level; j <= max_level; j++ {
        col_vals := make([]int, 0)
        for i := 0; i <= max_level; i++ {     
            level := Level{
                row: i,
                col: j,
            }

            if nodes, ok := hash_map[level]; ok {
                if len(nodes.data) > 1 {
                    sort.Ints(nodes.data)
                }

                col_vals = append(col_vals, nodes.data...)
            }
        }

        if len(col_vals) > 0 {
            result = append(result, col_vals)
        }
    }

    return result
}

func helper(root *TreeNode, row int, col int, hash_map map[Level]Nodes) int {
    if root == nil {
        return 0
    }

    level := Level{
        row: row, 
        col: col,
    }

    nodes, ok := hash_map[level]

    if !ok {
        hash_map[level] = Nodes{
            data: []int{root.Val},
        }
    } else {
        hash_map[level] = Nodes{
            data: append(nodes.data, root.Val),
        }
    }

    left := helper(root.Left, row + 1, col - 1, hash_map)
    right := helper(root.Right, row + 1, col + 1, hash_map)

    return 1 + max(left, right)
}