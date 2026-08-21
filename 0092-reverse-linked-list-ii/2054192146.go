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

/*
    index := 1 // treating as 1th index LL
    var start *ListNode
    var tail *ListNode
    var current *ListNode = head
    var next *ListNode = current.Next
    var first *ListNode

    for index <= right {
        if index <= left {
            start = tail
            tail = current
            current = current.Next
            if next != nil {
                next = next.Next
            }
        } else {
            if first == nil {
                start = tail
                tail = current
            }

            current.Next = first
            first = current
            current = next

            if next != nil {
                next = next.Next   
            }   
        }

        index++
    }

    tail.Next = next
    if start != nil {
        start.Next = current
    }

    return head
*/