/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    nodeToNew := make(map[*Node]*Node)

    for curr := head; curr != nil; curr = curr.Next {
        nodeToNew[curr] = &Node{Val: curr.Val}
    }

    for curr := head; curr != nil; curr = curr.Next {
        nodeToNew[curr].Next = nodeToNew[curr.Next] 
        nodeToNew[curr].Random = nodeToNew[curr.Random]
    }

    return nodeToNew[head]
}