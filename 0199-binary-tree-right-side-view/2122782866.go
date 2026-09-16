/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    result := make([]int, 0)

    if root == nil {
        return result
    }

    helper(root, 0, &result)
    return result
}

func helper(root *TreeNode, level int, result *[]int) {
    if root == nil {
        return 
    }

    if level == len(*result){
        *result = append(*result, root.Val)
    }

    helper(root.Right, level + 1, result)
    helper(root.Left, level + 1, result)
}
