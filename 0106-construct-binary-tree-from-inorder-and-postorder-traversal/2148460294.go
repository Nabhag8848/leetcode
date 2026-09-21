/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    inorder_idx := make(map[int]int, len(inorder))

    for idx := range inorder {
        inorder_idx[inorder[idx]] = idx
    }
    
    return helper(postorder, 0, len(postorder) - 1, inorder, 0, len(inorder) - 1, inorder_idx)
}

func helper(postorder []int, post_start, post_end int, inorder []int, in_start, in_end int, inorder_idx map[int]int) *TreeNode {
    if post_start > post_end || in_start > in_end {
        return nil
    }

    root_idx := inorder_idx[postorder[post_end]]
    root_val := postorder[post_end]
    root := &TreeNode{
        Val: root_val,
        Left: nil,
        Right: nil,
    }

    nums_left := root_idx - in_start

    root.Left = helper(postorder, post_start, post_start + nums_left - 1, inorder, in_start, root_idx - 1, inorder_idx)
    root.Right = helper(postorder, post_start + nums_left, post_end - 1, inorder, root_idx + 1, in_end, inorder_idx)
    return root
}