/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
    result := make([][]int, 0)
    arr := make([]int, 0)
    inorder(root, &arr)

    for _,query := range queries {
        min_max := []int{-1, -1}
        idx := sort.SearchInts(arr, query)

        if idx > 0 {
            min_max[0] = arr[idx - 1]
            if arr[idx - 1] == query {
                min_max[0] = query
                min_max[1] = query
                continue
            }
        }

        if idx < len(arr)  {  
            if arr[idx] == query {
                min_max[0] = query
                min_max[1] = query
            } else {
                min_max[1] = arr[idx]
            }
        }

        result = append(result, min_max)
    }

    return result
}

func inorder(node *TreeNode, vals *[]int) {
    if node == nil {
        return
    }
    inorder(node.Left, vals)
    *vals = append(*vals, node.Val)
    inorder(node.Right, vals)
}
/*
   query = 5

   root 
   > smallest_max = root can be root but we need to find on left. largest_min = -1 (find that on left)
   < largest_min = root, find it on right. smallest_max = -1 find on right
   = smallest_max = root. we are sure. largest_min = root
*/