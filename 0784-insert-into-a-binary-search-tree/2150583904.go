/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    node := &TreeNode{
        Val: val, 
        Left: nil,
        Right: nil,
    }

    if root == nil {
        return node
    }

    var curr *TreeNode = root

    for curr != nil {
        if curr.Val > val {
            if curr.Left == nil {
                curr.Left = node
                break
            }

            curr = curr.Left
        } else {
            if curr.Right == nil {
                curr.Right = node
                break
            }
            curr = curr.Right
        }
    }

    return root
}