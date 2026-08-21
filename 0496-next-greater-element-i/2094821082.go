func nextGreaterElement(nums1 []int, nums2 []int) []int {
    st := NewStack[int]()
    hash_map := make(map[int]int)
    ans := make([]int, len(nums1))

    for i := len(nums2) - 1; i >= 0; i-- {
        if st.Empty() {
            hash_map[nums2[i]] = -1
        } else if nums2[i] < st.Peek() {
            hash_map[nums2[i]] = st.Peek()
        } else {
            for !st.Empty() && nums2[i] > st.Peek() {
                st.Pop()
            }

            if st.Empty() {
                hash_map[nums2[i]] = -1
            } else {
                hash_map[nums2[i]] = st.Peek()
            }
        }

        st.Push(nums2[i])
    }

    for i := range nums1 {
        value,_ := hash_map[nums1[i]]
        ans[i] = value
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