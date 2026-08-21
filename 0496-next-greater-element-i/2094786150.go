func nextGreaterElement(nums1 []int, nums2 []int) []int {
    s1 := NewStack[int]()
    s2 := NewStack[int]()

    ans := make([]int, len(nums1))

    for i := range nums2 {
        s1.Push(nums2[len(nums2) - (i + 1)])
    }

    for i := range nums1 {
        for !s1.Empty() && s1.Peek() != nums1[i] {
            s2.Push(s1.Pop())
        }

        for !s1.Empty() && s1.Peek() <= nums1[i] {
            s2.Push(s1.Pop())   
        }


        if s1.Empty() {
            ans[i] = -1
        } else {
            ans[i] = s1.Peek()
        }

        for !s2.Empty() {
            s1.Push(s2.Pop())
        }
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