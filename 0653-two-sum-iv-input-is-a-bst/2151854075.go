/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    hash_map := make(map[int]int)
    inorder(root, hash_map)
    for y, count := range hash_map {
        x := k - y
        if _, ok := hash_map[x]; ok {
            if x != y {
                return true
            }

            if count > 1 {
                return true
            }
        } 
    }

    return false
}

func inorder(node *TreeNode, hash_map map[int]int) {
    if node == nil {
        return
    }
    inorder(node.Left, hash_map)
    hash_map[node.Val]++
    inorder(node.Right, hash_map)
}