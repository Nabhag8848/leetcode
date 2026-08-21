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

    tail := head
    numNode := len - rotate
    index := 1
    var newTail *ListNode = head


    for tail.Next != nil {
        tail = tail.Next

        if index < numNode {
            newTail = newTail.Next
            index++ 
        }
    }

    tail.Next = head
    head = newTail.Next
    newTail.Next = nil

    return head
}

func lengthOfLL(head *ListNode) int {
    count := 0

    for head != nil {
        head = head.Next
        count++
    }

    return count
}