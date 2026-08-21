/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    hash_map := make(map[*ListNode]struct{})

    currA := headA
    currB := headB

    for currA != nil  {
        hash_map[currA] = struct{}{}
        currA = currA.Next
    }

    for currB != nil {
        if _,is_exist := hash_map[currB]; is_exist {
            return currB
        }

        currB = currB.Next 
           
    }
    
    return nil  
}