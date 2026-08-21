/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    first := head
    second := head.Next

    for second != nil {
        for second != nil && first.Val == second.Val {
            temp := second
            second = second.Next
            temp.Next = nil
        }

        first.Next = second
        first = second

        if second != nil {
            second = second.Next
        }
    }

    return head
}