/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prevGroupTail := dummy

	for {
		kth := prevGroupTail
		for i := 0; i < k && kth != nil; i++ {
			kth = kth.Next
		}
		if kth == nil {
			break
		}

		groupNext := kth.Next
		prev := groupNext
		curr := prevGroupTail.Next

		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		oldGroupHead := prevGroupTail.Next
		prevGroupTail.Next = kth
		prevGroupTail = oldGroupHead
	}

	return dummy.Next
}
