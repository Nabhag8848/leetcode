/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
    value := root.Val
    left := root.Left.Val 
    right := root.Right.Val

    return  value == (left + right)
}