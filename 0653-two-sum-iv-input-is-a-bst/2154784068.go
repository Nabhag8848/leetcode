/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    var lower *BSTIterator = NewBSTIterator(root, false)
    var upper *BSTIterator = NewBSTIterator(root, true)

    left := lower.Next()
    right := upper.Next()

    for left < right {
        fmt.Println(left, right)
        sum := left + right

        if sum == k {
            return true
        }

        if sum < k {
            left = lower.Next()
        } else {
            right = upper.Next()
        }
    }

    return false
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    st *Stack[*TreeNode]
    decreasing bool
}


func NewBSTIterator(root *TreeNode, decreasing bool) *BSTIterator {
    st := NewStack[*TreeNode]()

    var curr *TreeNode = root

    if decreasing {
        for curr != nil {
            st.Push(curr)
            curr = curr.Right
        }
    } else {
        for curr != nil {
            st.Push(curr)
            curr = curr.Left
        }
    }


    return &BSTIterator{
        st: st,
        decreasing: decreasing,
    }
}


func (this *BSTIterator) Next() int {
    top := this.st.Pop()
    if this.decreasing {
        left := top.Left
        for left != nil {
            this.st.Push(left)
            left = left.Right
        }
    } else {
        right := top.Right
        for right != nil {
            this.st.Push(right)
            right = right.Left
        }
    }
   
    return top.Val
}


func (this *BSTIterator) HasNext() bool {
    return !this.st.Empty()
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


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */