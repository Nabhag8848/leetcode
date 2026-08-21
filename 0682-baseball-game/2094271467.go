func calPoints(operations []string) int {
    st := NewStack[int]()

    for i, _ := range operations {
        if operations[i] == "C" {
            if !st.Empty() {
                st.Pop()
            }
            
        } else if operations[i] == "D" {
            num := st.Peek() * 2
            st.Push(num)

        } else if operations[i] == "+" {
            num2 := st.Pop()
            num1 := st.Peek()
            st.Push(num2)

            st.Push(num1 + num2)
        } else {
            num,_ := strconv.Atoi(operations[i])
            st.Push(num)
        }
    }

    sum := 0

    for !st.Empty() {
        sum += st.Pop()
    }

    return sum
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