/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root, target *TreeNode, k int) []int {
    result := make([]int, 0)
    path := make([]*TreeNode, 0)

    if !findPath(root, target, &path) {
        return result
    }

    var blocked *TreeNode
    distance := 0

    for i := len(path) - 1; i >= 0 && distance <= k; i-- {
        collect(path[i], blocked, k-distance, &result)
        blocked = path[i]
        distance++
    }

    return result
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

func collect(node, blocked *TreeNode, remaining int, result *[]int) {
    if node == nil || node == blocked || remaining < 0 {
        return
    }

    if remaining == 0 {
        *result = append(*result, node.Val)
        return
    }

    collect(node.Left, blocked, remaining-1, result)
    collect(node.Right, blocked, remaining-1, result)
}