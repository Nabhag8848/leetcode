/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    curr1 := l1
    curr2 := l2
    var head *ListNode
    var tail *ListNode
    var carry int = 0

    for curr1 != nil || curr2 != nil {
        sum := carry

        if curr1 != nil {
            sum = sum + curr1.Val
        }

        if curr2 != nil {
            sum = sum + curr2.Val
        }

        var digit int

        if sum > 9 {
            digit = sum % 10
            carry = sum / 10
        } else {
            digit = sum
            carry = 0
        }

        node := &ListNode{
            Val: digit,
            Next: nil,
        }

        if head == nil {
            head = node
            tail = node
        } else {
            tail.Next = node
            tail = tail.Next
        }  

        if curr1 != nil {
            curr1 = curr1.Next
        } 

        if curr2 != nil {
            curr2 = curr2.Next 
        }    
    }

    if carry > 0 {
        node := &ListNode{
            Val: carry,
            Next: nil,
        }

        tail.Next = node
    }

    return head
}
