/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) int {
    sum := 0
    helper(root, &sum)
    return sum
}

func helper(root *TreeNode, sum *int) (int, int, int) {
    if root == nil {
        return math.MaxInt64, math.MinInt64, 0
    }

    left_min, left_max, left_sum := helper(root.Left, sum)
    right_min, right_max, right_sum := helper(root.Right, sum)
    total := left_sum + right_sum + root.Val
    if left_max < root.Val && right_min > root.Val {

        if *sum < total {
            *sum = total
        }

        return min(left_min, root.Val), max(right_max, root.Val), total
    }

    return math.MinInt64, math.MaxInt64, 0
}