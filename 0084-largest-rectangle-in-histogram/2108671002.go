func largestRectangleArea(heights []int) int {
    nse := findNextSmallerElement(heights)
    pse := findPrevSmallerElement(heights)
    result := 0

    for i := range heights {
        area := (nse[i] - pse[i] - 1) * heights[i]
        result = max(area, result)
    }

    return result
}


func findPrevSmallerElement(heights []int) []int{
    result := make([]int, len(heights))
    st := NewStack[int]()

    for i := range heights {
        for !st.Empty() && heights[st.Peek()] >= heights[i] {
            st.Pop()
        }

        if st.Empty() {
            result[i] = -1
        } else {
            result[i] = st.Peek()
        }

        st.Push(i)
    }

    return result
}


func findNextSmallerElement(heights []int) []int{
    result := make([]int, len(heights))
    st := NewStack[int]()

    for i := len(heights) - 1; i >= 0; i-- {
        for !st.Empty() && heights[st.Peek()] >= heights[i] {
            st.Pop()
        }

        if st.Empty() {
            result[i] = len(heights)
        } else {
            result[i] = st.Peek()
        }

        st.Push(i)
    }

    return result
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

/*
    5,5,6,2,6

    1 2 3 4 5 

    5 4 3 2 1

*/ 