func removeKdigits(num string, k int) string {
    st := NewStack[byte]()

    for i := range num {
        for !st.Empty() && st.Peek() > num[i] && k > 0 {
            k--
            st.Pop()
        }

        st.Push(num[i])
    }

    for k > 0 {
        st.Pop()
        k--
    }

    if st.Empty() {
        return "0"
    }
    
    var builder strings.Builder

    tmp := NewStack[byte]()
    
    for !st.Empty() {
        tmp.Push(st.Pop())
    }

    for !tmp.Empty() && tmp.Peek() == '0' {
        tmp.Pop()
    }

    for !tmp.Empty() {
        builder.WriteByte(tmp.Pop())
    }

    if builder.Len() == 0 {
        return "0"
    }

    return builder.String()
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
    1 4 3 2 2 1 9

    <start_min> <smallest possible> <smallest_possible>


    1 0 2 0 0


    1430232
*/