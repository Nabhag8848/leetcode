/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    arr := make([]int, 0)
    i := 0
    inorder(root, &arr, false, &i)
    sort.Ints(arr)
    inorder(root, &arr, true, &i)
}

func inorder(node *TreeNode, vals *[]int, assign bool, idx *int) {
    if node == nil {
        return
    }
    inorder(node.Left, vals, assign, idx)
    if assign {
        node.Val = (*vals)[*idx]
        *idx = *idx + 1
    } else {
        *vals = append(*vals, node.Val)
    }
    inorder(node.Right, vals, assign, idx)
}