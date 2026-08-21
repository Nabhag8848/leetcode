/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    middle := middleNode(head)
    leftHead := head
    rightHead := middle.Next
    middle.Next = nil

    left := sortList(leftHead)
    right := sortList(rightHead)

    return mergeTwoLists(left, right)
}


func middleNode(head *ListNode) *ListNode {
    fast := head
    slow := head
    var prev *ListNode

    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    return prev
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head *ListNode
    var tail *ListNode
    var node *ListNode

    for list1 != nil && list2 != nil {
        if list1.Val > list2.Val {
            node = list2
            list2 = list2.Next
            node.Next = nil
        } else if list1.Val < list2.Val {
            node = list1
            list1 = list1.Next
            node.Next = nil
        } else {
            node = list2
            list2 = list2.Next

            node.Next = list1
            list1 = list1.Next

            if node != nil && node.Next != nil {
                node.Next.Next = nil
            }
           
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