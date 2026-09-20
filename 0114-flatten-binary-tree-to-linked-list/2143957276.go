/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
    var prev *TreeNode

    var helper func(*TreeNode)
    helper = func(node *TreeNode) {
        if node == nil {
            return
        }

        helper(node.Right)
        helper(node.Left)

        node.Right = prev
        node.Left = nil
        prev = node
    }

    helper(root)
}
