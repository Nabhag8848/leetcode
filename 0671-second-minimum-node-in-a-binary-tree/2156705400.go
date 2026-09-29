/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    mini := math.MaxInt64
    helper(root, root.Val, &mini)

    if mini == math.MaxInt64 {
        return -1
    }

    return mini
}

func helper(root *TreeNode, lower int, mini *int) {
    if root == nil {
        return 
    }

    if root.Val > lower && *mini > root.Val {
        *mini = root.Val
    }

    helper(root.Left, lower, mini)
    helper(root.Right, lower, mini)
}