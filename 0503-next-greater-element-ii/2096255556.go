func nextGreaterElements(nums []int) []int {
    st := NewStack[int]()
    ans := make([]int, len(nums))
    length := len(nums)

    for i := 2 * length - 1; i >= 0; i-- {
        for !st.Empty() && nums[i % length] >= st.Peek() {
            st.Pop()
        }

        if i < length {
            if st.Empty(){
                ans[i] = -1
            } else {
                ans[i] = st.Peek()
            } 
        }


        st.Push(nums[i % length])
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