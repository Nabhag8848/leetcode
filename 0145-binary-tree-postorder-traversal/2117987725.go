/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    arr := make([]int, 0)
    helper(root, &arr)
    return arr
}

func helper(node *TreeNode, arr *[]int) {
    if node == nil {
        return
    }
    
    helper(node.Left, arr)
    helper(node.Right, arr)
    *arr = append(*arr, node.Val)
}