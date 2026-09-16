/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
   return helper(root, "")
}

func helper(root *TreeNode, str string) []string {
    var arr []string
    str = str + strconv.Itoa(root.Val)

    if root.Left != nil {
        arr = append(arr, helper(root.Left, str + "->")...)
    }

    if root.Right != nil {
        arr = append(arr, helper(root.Right, str + "->")...)
    }

    if root.Right == nil && root.Left == nil {
        arr = append(arr, str)
    }

    return arr
}
