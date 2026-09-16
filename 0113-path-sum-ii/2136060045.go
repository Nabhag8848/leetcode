/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    arr := make([][]int, 0)
    helper(root, targetSum, &arr, []int{})
    return arr
}

func helper(root *TreeNode, targetSum int, arr *[][]int, res []int) bool {
    if root == nil {
        return false
    }

    targetSum = targetSum - root.Val
    res = append(res, root.Val)

    if root.Left == nil && root.Right == nil {
        if targetSum == 0 {
            final := make([]int, len(res))
            copy(final, res)
            *arr = append(*arr, final)
            
            return true
        }

        return false
    }

    left := helper(root.Left, targetSum, arr, res)
    right := helper(root.Right, targetSum, arr,res)
 
    return left || right
}