/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    nodeToRandom := make(map[*Node]*Node)
    nodeToNew := make(map[*Node]*Node)
    var curr *Node = head
    var newHead *Node
    var tail *Node

    for curr != nil {
        nodeToRandom[curr] = curr.Random
        curr = curr.Next
    }

    curr = head

    for curr != nil {
        var node *Node
        if newHead == nil {
            node = &Node{
                Val:curr.Val,
                Next: nil,
                Random:nil,
            }
            newHead = node
            tail = node
        } else {
            node = &Node{
                Val:curr.Val,
                Next:nil,
                Random:nil,
            }

            tail.Next = node
            tail = node
        }

        nodeToNew[curr] = node
        curr = curr.Next  
    }

    curr = head

    for curr != nil {
        var currNew *Node = nodeToNew[curr]
        var currRandom *Node = nodeToRandom[curr]

        var currNewRandom *Node = nodeToNew[currRandom]
        currNew.Random = currNewRandom
        
        curr = curr.Next
    }

    return newHead
}