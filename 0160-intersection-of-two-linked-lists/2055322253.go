/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    currA := headA
    currB := headB

    if currA == nil || currB == nil {
        return nil
    } 

    for currA != currB {
        currA = currA.Next
        currB = currB.Next

        if currA == currB {
            return currA
        }

        if currA == nil {
            currA = headB
        }
        
        if currB == nil {
            currB = headA
        }

    }

    return currA
}