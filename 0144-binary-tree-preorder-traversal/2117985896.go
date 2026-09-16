/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    arr := make([]int, 0)
    helper(root, &arr)
    return arr
}

func helper(node *TreeNode, arr *[]int){
    if node == nil {
        return
    }
    
    *arr = append(*arr, node.Val)

    helper(node.Left, arr)
    helper(node.Right, arr)
}