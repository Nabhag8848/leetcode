/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    var head *ListNode
    var tail *ListNode
    var mini int = math.MaxInt32

    for {
        for i := range lists {
            if lists[i] != nil {
                mini = min(lists[i].Val, mini)
            }
        }

        if mini == math.MaxInt32 {
            break
        }

        for i := range lists {
            if lists[i] != nil && lists[i].Val == mini {
                if head == nil {
                    head = lists[i]
                    tail = lists[i]

                    new := lists[i].Next
                    lists[i].Next = nil
                    lists[i] = new
                } else {
                    tail.Next = lists[i]
                    tail = lists[i]
                    new := lists[i].Next
                    lists[i].Next = nil
                    lists[i] = new
                }
            }
        }

        mini = math.MaxInt32
    }

    return head
}