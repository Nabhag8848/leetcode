/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if result := helper(root); result == -1 {
        return false
    }

    return true
}

func helper(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := helper(root.Left)
    if left == -1 {
        return left
    }

    right := helper(root.Right)

    if right == -1 {
        return right
    }

    if abs(left - right) > 1 {
        return -1
    }

    return 1 + max(left, right)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
