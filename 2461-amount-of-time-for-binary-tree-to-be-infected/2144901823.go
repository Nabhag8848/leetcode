/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	target := findTarget(root, start)
	path := make([]*TreeNode, 0)
	findPath(root, target, &path)

	var blocked *TreeNode
	answer := 0

	for i := len(path) - 1; i >= 0; i-- {
		distanceFromStart := len(path) - 1 - i
		answer = max(answer, collect(path[i], blocked, distanceFromStart))
		blocked = path[i]
	}

	return answer
}

func findTarget(root *TreeNode, start int) *TreeNode {
	if root == nil || root.Val == start {
		return root
	}

	if left := findTarget(root.Left, start); left != nil {
		return left
	}

	return findTarget(root.Right, start)
}

func findPath(node, target *TreeNode, path *[]*TreeNode) bool {
	if node == nil {
		return false
	}

	*path = append(*path, node)

	if node == target ||
		findPath(node.Left, target, path) ||
		findPath(node.Right, target, path) {
		return true
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func collect(node, blocked *TreeNode, distance int) int {
	if node == nil || node == blocked {
		return distance - 1
	}

	left := collect(node.Left, blocked, distance+1)
	right := collect(node.Right, blocked, distance+1)

	return max(distance, max(left, right))
}
