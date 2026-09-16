/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var diameter int = 0
    helper(root, &diameter)
    return diameter
}

func helper(root *TreeNode, diameter *int) int {
    if root == nil {
        return 0
    }

    left := helper(root.Left, diameter)
    right := helper(root.Right, diameter)
    *diameter = max(*diameter, left + right)

    return 1 + max(left, right)
}