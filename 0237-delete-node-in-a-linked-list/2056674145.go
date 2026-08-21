/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    var prev *ListNode = node
    var head *ListNode = prev.Next

    for head.Next != nil {
        prev.Val = head.Val

        prev = head
        head = head.Next
    }

    prev.Val = head.Val
    prev.Next = nil
}