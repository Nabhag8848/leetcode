/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val == key {
        return helper(root)
    }

    var curr *TreeNode = root

    for curr != nil {
        if curr.Val > key {
            if curr.Left != nil && curr.Left.Val == key {
                curr.Left = helper(curr.Left)
                break
            }

            curr = curr.Left
        } else {
            if curr.Right != nil && curr.Right.Val == key {
                curr.Right = helper(curr.Right)
                break
            }

            curr = curr.Right
        }
    }

    return root
}

func helper(root *TreeNode) *TreeNode {
    if root.Left == nil {
        return root.Right
    }

    if root.Right == nil {
        return root.Left
    }

    rightChild := root.Right
    lastRight := findLastRight(root.Left)
    lastRight.Right = rightChild
    return root.Left
}

func findLastRight(root *TreeNode) *TreeNode {
    if root.Right == nil {
        return root
    }

    return findLastRight(root.Right)
}