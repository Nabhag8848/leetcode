/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    var i int = 0
    return helper(preorder, &i, math.MaxInt32)
}

func helper(preorder []int, idx *int, upper int) *TreeNode {
    if *idx >= len(preorder) {
        return nil
    }

    val := preorder[*idx]

    if val > upper {
        return nil
    }

    node := &TreeNode{
        Val: val,
    }

    *idx = *idx + 1
    node.Left = helper(preorder, idx, node.Val)
    node.Right = helper(preorder, idx, upper)

    return node
}