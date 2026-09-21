/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    inorder_idx := make(map[int]int, len(inorder))

    for idx := range inorder {
        inorder_idx[inorder[idx]] = idx
    }

    return helper(preorder, 0, len(preorder) - 1, inorder, 0, len(inorder) - 1, inorder_idx)
}

func helper(preorder []int, pre_start, pre_end int, inorder []int, in_start, in_end int, inorder_idx map[int]int) *TreeNode {
    if pre_start > pre_end || in_start > in_end {
        return nil
    }
    root_idx := inorder_idx[preorder[pre_start]]
    root_val := preorder[pre_start]
    root := &TreeNode{
        Val: root_val,
        Left: nil,
        Right: nil,
    }
    nums_left := root_idx - in_start

    root.Left = helper(preorder, pre_start + 1, pre_start + nums_left, inorder, in_start, root_idx - 1, inorder_idx)
    root.Right = helper(preorder, pre_start + nums_left + 1, pre_end, inorder, root_idx + 1, in_end, inorder_idx)

    return root
}