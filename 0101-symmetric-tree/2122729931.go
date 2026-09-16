/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return helper(root.Left, root.Right)
}

func helper(ltree *TreeNode, rtree *TreeNode) bool {
    if ltree == nil || rtree == nil {
        return ltree == rtree
    }

    if ltree.Val != rtree.Val {
        return false
    }

    left := helper(ltree.Left, rtree.Right)

    if !left {
        return false
    }

    right := helper(ltree.Right, rtree.Left)

    if !right {
        return false
    }

    return true
}