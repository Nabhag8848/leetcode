/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := leftHeight(root)
    right := rightHeight(root)

    if left == right {
        return (1 << left) - 1
    }

    return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func rightHeight(root *TreeNode) int {
    node := root
    count := 0
    for node != nil {
        node = node.Right
        count++
    }
    return count
}

func leftHeight(root *TreeNode) int {
    node := root
    count := 0
    for node != nil {
        node = node.Left
        count++
    }
    return count
}