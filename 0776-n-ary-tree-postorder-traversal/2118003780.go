/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    if root == nil {
        return []int{}
    }

    res := []int{}

    for _,node := range root.Children {  
        res = append(res, postorder(node)...)
    }

    res = append(res, root.Val)
    
    return res
}