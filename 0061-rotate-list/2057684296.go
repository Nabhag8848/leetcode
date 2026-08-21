/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    
    if head == nil {
        return head 
    }


    len := lengthOfLL(head)
    rotate := k % len

    if rotate == 0 {
        return head
    }

    head = reverseList(head)

    start := head
    curr := head
    index := 1

    for index < rotate {
        curr = curr.Next
        index++
    }

    head = curr.Next
    curr.Next = nil

    newHead := reverseList(start)
    start.Next = reverseList(head)

    return newHead
}

/*
    5 4 3 2 1

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

func lengthOfLL(head *ListNode) int {
    count := 0

    for head != nil {
        head = head.Next
        count++
    }

    return count
}