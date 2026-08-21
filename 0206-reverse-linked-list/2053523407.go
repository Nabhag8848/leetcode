/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    if head == nil {
        return head
    }

    var prev *ListNode
    var current *ListNode = head
    var next *ListNode = current.Next

    for next != nil {
        current.Next = prev
        prev = current
        current = next
        next = next.Next
    }

    current.Next = prev

    return current
}