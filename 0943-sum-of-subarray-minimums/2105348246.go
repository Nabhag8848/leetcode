func sumSubarrayMins(arr []int) int {
    total := 0
    nse := findNextSmallerElement(arr)
    pse := findPrevSmallerElement(arr)

    MOD := 1_000_000_007

    for i := range arr {
        left := i - pse[i]
        right := nse[i] - i

        total = total + ((left * right * arr[i]) % MOD)
    }

    return total % MOD
}

func findPrevSmallerElement(arr []int) []int {
    pse := make([]int, len(arr))
    st := NewStack[int]()

    for i := range arr {
        for !st.Empty() && arr[st.Peek()] > arr[i] {
            st.Pop()
        }

        if st.Empty() {
            pse[i] = -1
        } else {
            pse[i] = st.Peek()
        }

        st.Push(i)
    }

    return pse
}

func findNextSmallerElement(arr []int) []int {
    nse := make([]int, len(arr))
    st := NewStack[int]()

    for i := len(arr) - 1; i >= 0; i-- {
        for !st.Empty() && arr[st.Peek()] >= arr[i] {
            st.Pop()
        }

        if st.Empty() {
            nse[i] = len(arr)
        } else {
            nse[i] = st.Peek()
        }

        st.Push(i)
    }

    return nse
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