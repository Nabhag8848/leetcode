/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Frame struct {
    node *TreeNode
    visited bool
}

func diameterOfBinaryTree(root *TreeNode) int {
    var diameter int = 0
    height := make(map[*TreeNode]int, 0)
    stack := NewStack[*Frame]()

    stack.Push(&Frame{
        node: root,
        visited: false,
    })

    for !stack.Empty() {
        top := stack.Pop()
        node := top.node
        is_visited := top.visited
        
        if is_visited {
            left := height[node.Left]
            right := height[node.Right]

            if diameter < left + right {
                diameter = left + right
            }

            height[node] = 1 + max(left, right)
        } else {
            stack.Push(&Frame{
                node: node,
                visited: !is_visited,
            })

            if node.Right != nil {
                stack.Push(&Frame{
                    node: node.Right,
                    visited: false,
                })    
            }

            if node.Left != nil {
                stack.Push(&Frame{
                    node: node.Left,
                    visited: false,
                })
            }
        }
    }

    return diameter
}

type Stack[T any] struct {
	data []T
	top  int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0), top: -1}
}

func (s *Stack[T]) Push(value T) {
	s.top++
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() T {
	var zero T
	if s.top < 0 {
		return zero
	}

	value := s.data[s.top]
    s.data = s.data[:s.top]
	s.top--

	if s.top == -1 {
		s.data = make([]T, 0)
	}
	return value
}

func (s *Stack[T]) Peek() T {
	var zero T
	if s.top < 0 {
		return zero
	}

	return s.data[s.top]
}

func (s *Stack[T]) Empty() bool {
	return s.top == -1
}

func (s *Stack[T]) Size() int {
	return s.top + 1
}