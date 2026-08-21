/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    hash_map := make(map[*ListNode]struct{}, 0)
    current := head

    for current != nil && current.Next != nil {
        if _,is_exist := hash_map[current]; is_exist {
            return true
        }   
        hash_map[current] = struct{}{}
        current = current.Next
    }
    return false
}