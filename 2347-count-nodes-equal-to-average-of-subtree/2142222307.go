/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    var ans int = 0
    helper(root, &ans)
    return ans
}

func helper(root *TreeNode, ans *int) (int, int) {
    if root == nil {
        return 0, 0
    }

    lsum, ln := helper(root.Left, ans)
    rsum, rn := helper(root.Right, ans)
    val := root.Val 
    sum := lsum + rsum + val
    n := ln + rn + 1

    if (sum / n) == val {
        *ans = *ans + 1
    }

    return sum, n
}