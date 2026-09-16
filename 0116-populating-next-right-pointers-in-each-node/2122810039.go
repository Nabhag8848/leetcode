/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
        return root
    }

    if root.Left != nil {
        root.Left.Next = root.Right
        if root.Next != nil {
            root.Right.Next = root.Next.Left
        }
    }

    connect(root.Left)
    connect(root.Right)

    return root
}
