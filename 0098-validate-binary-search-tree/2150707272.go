/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
    if node == nil {
        return true
    }
    
    // Check if current node violates bounds
    if int64(node.Val) <= min || int64(node.Val) >= max {
        return false
    }
    
    // Validate left and right subtrees with updated bounds
    left := validate(node.Left, min, int64(node.Val))
    if !left {
        return false
    }
    
    right := validate(node.Right, int64(node.Val), max)
    if !right {
        return false
    }
    
    return true
}