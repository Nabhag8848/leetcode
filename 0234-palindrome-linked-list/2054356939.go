/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
   middle,len := middleNode(head)
   isOdd := len % 2 == 1
   if isOdd {
      middle = middle.Next
   }

   newHead := reverse(middle)

   for newHead != nil && head != nil {
        if head.Val != newHead.Val {
            return false
        }

        head = head.Next
        newHead = newHead.Next
   }

   return true
}

func middleNode(head *ListNode) (*ListNode,int) {
    fast := head
    slow := head 
    count := 0

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        count = count + 2
    }

    if fast != nil {
        count++
    }

    return slow,count
}

func reverse(head *ListNode) *ListNode {

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