/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    mid := middle(head)

    second := reverse(mid.Next)
    
    mid.Next = nil

    first := head

    for second != nil {
        firstNext := first.Next
        secondNext := second.Next

        first.Next = second
        second.Next = firstNext

        first = firstNext
        second = secondNext
    }
}

func middle(head *ListNode) *ListNode {
    slow, fast := head, head

    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev
}