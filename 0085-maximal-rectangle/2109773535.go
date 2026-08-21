func maximalRectangle(matrix [][]byte) int {
    arr := make([]int, len(matrix[0]))
    maxi := 0
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == '0' {
                arr[j] = 0
            } else {
                arr[j] = arr[j] + 1
            }
        }

        maxi = max(maxi, largestRectangleArea(arr))
    }

    return maxi
}

func largestRectangleArea(heights []int) int {
    st := NewStack[int]()
    result := 0

    for i := range heights {
        for !st.Empty() && heights[st.Peek()] > heights[i] {
            element := heights[st.Pop()]
            nse := i
            var pse int
            if ok := st.Empty(); ok {
                pse = -1
            } else {
                pse = st.Peek()
            }


            area := element * (nse - pse - 1)
            result = max(result, area)
        }

        st.Push(i)
    }

    for !st.Empty() {
        nse := len(heights)
        element := heights[st.Pop()]
        var pse int
        if ok := st.Empty(); ok {
            pse = -1
        } else {
            pse = st.Peek()
        }

        area := element * (nse - pse - 1)
        result = max(result, area)
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


*/