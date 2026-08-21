func evalRPN(tokens []string) int {
    st := NewStack[int]()

    for i := range tokens {
        if tokens[i] == "+" {
            num2 := st.Pop()
            num1 := st.Pop()

            st.Push(num2 + num1)
        } else if tokens[i] == "/"{
            num2 := st.Pop()
            num1 := st.Pop()

            st.Push(num1 / num2)

        } else if tokens[i] == "-" {
            num2 := st.Pop()
            num1 := st.Pop()

            st.Push(num1 - num2)

        } else if tokens[i] == "*" {
            num2 := st.Pop()
            num1 := st.Pop()

            st.Push(num1 * num2)

        } else {
            num,_ := strconv.Atoi(tokens[i])
            st.Push(num)
        }
    }

    return st.Pop()
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