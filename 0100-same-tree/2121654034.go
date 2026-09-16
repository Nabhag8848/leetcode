/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
        return p == q
    }

    if p.Val != q.Val {
        return false
    }

    left := isSameTree(p.Left, q.Left)

    if !left {
        return false
    }

    right := isSameTree(p.Right, q.Right)

    if !right {
        return false
    }

    return true
}

