/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    moves := 0
    helper(root, &moves)
    return moves
}

func helper(root *TreeNode, moves *int) int {
    if root == nil {
        return 0
    }

    left := helper(root.Left, moves)
    right := helper(root.Right, moves)
    
    *moves += abs(left) + abs(right)

    return root.Val + left + right - 1
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
