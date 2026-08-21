func trap(height []int) int {
    result := 0
    stack := NewStack[int]() 

    for i := len(height) - 1; i >= 0; i-- {
        for !stack.Empty() && height[i] > height[stack.Peek()] {
            top := stack.Pop()

            if stack.Empty() {
                break 
            }

            right := stack.Peek()
            width := right - i - 1
            boundedHeight := min(height[i], height[right]) - height[top]
            result += width * boundedHeight
        }
        stack.Push(i)
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
    0,1,0,2,1,0,1,3,2,1,2,1

    11
*/