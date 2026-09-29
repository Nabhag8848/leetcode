/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    mini := math.MaxInt32
    assign := false

    helper(root, root.Val, &mini, &assign)

    if assign {
        return mini
    }

    return -1
}

func helper(root *TreeNode, lower int, mini *int, assign *bool) {
    if root == nil {
        return 
    }

    if root.Val > lower && *mini >= root.Val {
        *assign = true
        *mini = root.Val
    }

    helper(root.Left, lower, mini, assign)
    helper(root.Right, lower, mini, assign)
}