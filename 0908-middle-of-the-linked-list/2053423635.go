/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    current := head
    len := lengthOfLL(current)
    mid := len / 2

    for i := 0; i < mid; i++ {
        head = head.Next
    }

    return head
}

func lengthOfLL(prev *ListNode) int {
    count := 0

    for prev != nil {
        count++
        prev = prev.Next
    }


    return count
}