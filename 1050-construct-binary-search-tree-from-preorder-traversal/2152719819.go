/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    min := math.MinInt32
    max := math.MaxInt32

    root := &TreeNode{
        Val: preorder[0],
    }

    parent := make(map[*TreeNode]*TreeNode, 0)
    min_max := make(map[*TreeNode][]int)
    min_max[root] = []int{min, max}
    curr := root

    for i := 1; i < len(preorder); i++ {
        val := preorder[i]
        child := &TreeNode{
            Val: val,
        }
        lower := min_max[curr][0]
        upper := min_max[curr][1]

        if val > lower && val < upper  {
            if curr.Val > val {
                curr.Left = child
                min_max[child] = []int{lower, curr.Val}
            } else {
                curr.Right = child
                min_max[child] = []int{curr.Val, upper}
            }
            parent[child] = curr
            curr = child
        } else {
            curr = parent[curr]
            i--
        }
        
    }

    return root
}