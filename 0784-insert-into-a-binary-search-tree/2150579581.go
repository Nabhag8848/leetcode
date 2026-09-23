/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{
            Val: val,
            Left: nil, 
            Right: nil,
        }
    }

    if root.Val > val {
        left := insertIntoBST(root.Left, val)
        root.Left = left
    } else {
        right := insertIntoBST(root.Right, val)
        root.Right = right
    }

    return root
}