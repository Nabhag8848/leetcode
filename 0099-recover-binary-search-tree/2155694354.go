/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type State struct {
	prev   *TreeNode
	first  *TreeNode
	middle *TreeNode
	last   *TreeNode
}

func recoverTree(root *TreeNode)  {
    var state *State = &State{
        prev: &TreeNode{  
            Val: math.MinInt32,
        },
    }

    inorder(root, state)
 
    if state.first != nil && state.middle != nil {
        if state.last == nil {
            state.first.Val, state.middle.Val = state.middle.Val, state.first.Val
        } else {
            state.first.Val, state.last.Val = state.last.Val, state.first.Val
        }
    }
   

}

func inorder(node *TreeNode, state *State) {
    if node == nil {
        return
    }
    inorder(node.Left, state)
    if state.prev != nil && state.prev.Val > node.Val {
        if state.first == nil {
            state.first = state.prev
            state.middle = node
        } else {
            state.last = node
        }
    }
    state.prev = node
    inorder(node.Right, state)
}