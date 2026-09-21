/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	curr := root

	for curr != nil {
		if curr.Right == nil {
			result = append(result, curr.Val)
			curr = curr.Left
		} else {
			prev := curr.Right

			for prev.Left != nil && prev.Left != curr {
				prev = prev.Left
			}

			if prev.Left == nil {
				prev.Left = curr
				result = append(result, curr.Val)
				curr = curr.Right
			} else {
				prev.Left = nil
				curr = curr.Left
			}
		}
	}

	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return result
}