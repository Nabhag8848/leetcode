/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
        return nil
    }

    if root == p || root == q {
        return root
    }

    var left *TreeNode
    var right *TreeNode
    
    if p.Val < root.Val || q.Val < root.Val {
        left = lowestCommonAncestor(root.Left, p, q)
    }
    
    if p.Val > root.Val || q.Val > root.Val {
        right = lowestCommonAncestor(root.Right, p, q)
    }

    if left == nil {
        return right
    }

    if right == nil {
        return left
    }

    return root
}