func secondGreaterElement(nums []int) []int {
    st1 := NewStack[int]()
    st2 := NewStack[int]()

    ans := make([]int, len(nums))

    for i := range nums {
        for !st2.Empty() && nums[i] > nums[st2.Peek()] {
            ans[st2.Peek()] = nums[i]
            st2.Pop()
        }

        tmp := NewStack[int]()
        for !st1.Empty() && nums[i] > nums[st1.Peek()] {
            tmp.Push(st1.Pop())
        }

        for !tmp.Empty() {
            st2.Push(tmp.Pop())
        }

        st1.Push(i)
    }

    for !st1.Empty() {
        ans[st1.Pop()] = -1
    }

    for !st2.Empty() {
        ans[st2.Pop()] = -1
    }
    
    return ans
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