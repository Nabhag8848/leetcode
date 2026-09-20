/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
    var curr *TreeNode = root

    for curr != nil {
        if curr.Left != nil {
            node := curr.Left
            for node.Right != nil {
                node = node.Right
            }

            node.Right = curr.Right
            
            curr.Right = curr.Left
            curr.Left = nil
        }   

        curr = curr.Right
    }
}
