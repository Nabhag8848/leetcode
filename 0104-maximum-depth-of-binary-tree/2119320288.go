/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    level := 1
    left := level + maxDepth(root.Left)
    right := level + maxDepth(root.Right)

    return max(left, right)
}