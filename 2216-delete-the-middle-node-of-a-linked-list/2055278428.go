/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteMiddle(head *ListNode) *ListNode {
    var prev *ListNode
    var slow *ListNode = head
    var fast *ListNode = head
    var next *ListNode = head.Next

    if next == nil {
        return nil
    }

    middleNode(&slow, &fast, &prev, &next)

    prev.Next = next
    slow.Next = nil

    return head
}

func middleNode(slow **ListNode, fast **ListNode, prev **ListNode, next **ListNode) {
    for *fast != nil && (*fast).Next != nil {
        *prev = *slow
        *slow = (*slow).Next
        *next = (*next).Next
        *fast = (*fast).Next.Next
    }
}