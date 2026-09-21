/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    curr := root

    for curr != nil {
        if curr.Left == nil {
            result = append(result, curr.Val)
            curr = curr.Right
        } else {
            prev := curr.Left

            for prev.Right != nil && prev.Right != curr {
                prev = prev.Right
            }

            if prev.Right == nil {
                prev.Right = curr
                curr = curr.Left
            } else {
                prev.Right = nil
                result = append(result, curr.Val)
                curr = curr.Right
            }
        }
    }

    return result
}