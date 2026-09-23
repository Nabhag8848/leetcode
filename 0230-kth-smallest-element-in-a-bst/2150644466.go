/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    return morrisInorder(root, k)
}

func morrisInorder(root *TreeNode, k int) int {
    curr := root
    count := 0

    for curr != nil {
        if curr.Left == nil {
            count++
            if count == k {
                return curr.Val
            }

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
                count++
                if count == k {
                    return curr.Val
                }
                curr = curr.Right
            }
        }
    }

    return -1
}