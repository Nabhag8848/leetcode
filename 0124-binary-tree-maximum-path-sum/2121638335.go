/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxi := math.MinInt32
    helper(root, &maxi)
    return maxi
}

func helper(root *TreeNode, maxi *int) int {
    if root == nil {
        return 0
    }

    left := max(0, helper(root.Left, maxi))
    right := max(0, helper(root.Right, maxi))

    *maxi = max(*maxi, root.Val + left + right)

    return root.Val + max(left, right)
}