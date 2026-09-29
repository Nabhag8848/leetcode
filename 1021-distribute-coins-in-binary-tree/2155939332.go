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

func helper(root *TreeNode, moves *int) {
    if root == nil {
        return
    }

    helper(root.Left, moves)
    helper(root.Right, moves)

    if root.Left != nil {
        coins := root.Left.Val - 1

        if coins >= 0 {
            root.Val += coins
            *moves += coins
        } else {
            root.Val += coins
            *moves -= coins
        }

        root.Left.Val = 1
    }

    if root.Right != nil {
        coins := root.Right.Val - 1

        if coins >= 0 {
            root.Val += coins
            *moves += coins
        } else {
            root.Val += coins
            *moves -= coins
        }

        root.Right.Val = 1
    }
}
