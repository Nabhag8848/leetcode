/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head *ListNode
    var tail *ListNode
    var node *ListNode

    for list1 != nil && list2 != nil {
        if list1.Val > list2.Val {
            node = &ListNode {
                Val: list2.Val,
                Next: nil,
            }

            list2 = list2.Next
        } else if list1.Val < list2.Val {
            node = &ListNode {
                Val: list1.Val,
                Next: nil,
            }
            list1 = list1.Next
        } else {
            node = &ListNode {
                Val: list2.Val,
                Next: &ListNode {
                    Val: list1.Val,
                    Next: nil,
                },
            }
           
            list1 = list1.Next
            list2 = list2.Next
        }


        if head == nil {
            head = node
        } else {
            tail.Next = node
        }
            
        if node.Next != nil {
            tail = node.Next
        } else {
            tail = node
        }
    }

    for list1 != nil {
        if tail != nil {
            tail.Next = list1
        } else {
            head = list1
        }
        tail = list1
        list1 = list1.Next
    }

    for list2 != nil {
        if tail != nil {
            tail.Next = list2
        } else {
            head = list2
        }
        tail = list2
        list2 = list2.Next
    }

    return head
}