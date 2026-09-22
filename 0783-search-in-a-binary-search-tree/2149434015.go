/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    for root != nil && root.Val != val {
        if root.Val > val {
            root = root.Left
        } else {
            root = root.Right
        }
    }

    return root
}