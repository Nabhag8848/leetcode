/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func widthOfBinaryTree(root *TreeNode) int {
    left_most := make(map[int]int,0)
    max_width := 0

    if root == nil {
        return max_width
    }

    var helper func(node *TreeNode, depth int, idx int)
    helper = func (node *TreeNode, depth int, idx int) {
        if node == nil {
            return
        }

        if _,seen := left_most[depth]; !seen {
            left_most[depth] = idx
        }

        max_width = max(max_width, idx - left_most[depth] + 1)

        helper(node.Left, depth + 1, 2 * idx)
        helper(node.Right, depth + 1, 2 * idx + 1)
    }

    helper(root, 0, 0)

    return max_width
}

