/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    index := 1
    var start *ListNode 
    var current *ListNode = head
    var next *ListNode = current.Next
    var prev *ListNode = nil
    var tail *ListNode = nil

    for index <= right {
        if index < left {
            start = current
            current = current.Next
            tail = current
            next = next.Next
            
        } else {
            current.Next = prev
            prev = current
            current = next
            if next != nil {
                next = next.Next
            }
        }

      
        index++
    }

    if tail != nil {
        tail.Next = current
    } else {
        head.Next = current
        head = prev
    }

    if start != nil {
        start.Next = prev
    }

    return head
}