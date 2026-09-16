/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

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
        if root.Right != nil {
            root.Left.Next = root.Right
        } else {
            target := root
            for target != nil && target.Next != nil {
                if target.Next.Left != nil {
                    root.Left.Next = target.Next.Left
                    break
                }
                if target.Next.Right != nil {
                    root.Left.Next = target.Next.Right
                    break
                }
                target = target.Next
            }

        }
    }

    if root.Right != nil {
        target := root
        for target != nil && target.Next != nil {
            if target.Next.Left != nil {
                root.Right.Next = target.Next.Left
                break
            }
            if target.Next.Right != nil {
                root.Right.Next = target.Next.Right
                break
            }

            target = target.Next
        }
    }


    connect(root.Right)
    connect(root.Left)

    return root
}


