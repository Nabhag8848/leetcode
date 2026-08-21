/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    head = reverseList(head)
    curr1 := head
    carry := 0

    for curr1 != nil {
        double := curr1.Val * 2 + carry
        digit := double % 10
        carry = double / 10 

        curr1.Val = digit
        curr1 = curr1.Next
    }

    head = reverseList(head)

    if carry > 0 {
        node := &ListNode{
            Val: carry,
            Next: head,
        }

        head = node
    }

    return head
}

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