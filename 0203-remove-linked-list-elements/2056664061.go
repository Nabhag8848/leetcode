/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }

    var prev *ListNode
    var curr *ListNode = head
    var next *ListNode = curr.Next

    for curr != nil {
        if curr.Val == val {
            if curr == head {
                head = head.Next
                prev = curr
                curr = next
                if curr != nil {
                    next = next.Next
                }
            } else {
                prev.Next = next
                curr = next
                if curr != nil {
                    next = next.Next
                }  
            }
        } else {
            prev = curr

            if curr != nil {
                curr = curr.Next
            }

            if next != nil {
                next = next.Next
            }
        }
    }

    return head

}