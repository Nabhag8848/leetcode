/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    len := lengthOfLL(head)

    if len - n == 0 {
        return head.Next
    }

    var prev *ListNode 
    var current *ListNode = head
    var next *ListNode = current.Next

    for i:=1; i <= len - n; i++ {
        prev = current
        current = current.Next
        next = next.Next
    }

    prev.Next = next
    current.Next = nil

    return head
}

func lengthOfLL(head *ListNode) int {
    current := head
    count := 0

    for current != nil {
        count++
        current = current.Next
    }

    return count
}