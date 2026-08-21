/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


/*.  nil -> head

    len(1) -> head
    len(2) -> head
    len(3)
*/

func oddEvenList(head *ListNode) *ListNode {
    var prev *ListNode = head

    if prev == nil {
        return head
    }

    var current *ListNode = prev.Next

    if current == nil {
        return head
    }
    
    var next *ListNode = current.Next
    var start *ListNode = current

    for next != nil {
        prev.Next = next
        current.Next = next.Next

        prev = next

        if prev != nil {
            current = prev.Next

            if current != nil {
                next = current.Next
            } else {
                next = nil
            }
        }
    }

    prev.Next = start

    return head
}

