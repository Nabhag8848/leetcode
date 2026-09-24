/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    arr := make([]int, 0)
    inorder(root, &arr)
    l := 0
    r := len(arr) - 1 

    for l < r {
        val := arr[l] + arr[r]
        if val > k {
            r--
        } else if val < k {
            l++
        } else {
            return true
        }
    }

    return false
}

func inorder(node *TreeNode, arr *[]int) {
    if node == nil {
        return
    }
    inorder(node.Left, arr)
    *arr = append(*arr, node.Val)
    inorder(node.Right, arr)
}