func asteroidCollision(asteroids []int) []int {
    st := NewStack[int]()
    arr := []int{}

    for i := range asteroids {
       if asteroids[i] < 0 {
          if st.Empty() {
            arr = append(arr, asteroids[i])
          } else {
            for !st.Empty() && int(math.Abs(float64(asteroids[i]))) > st.Peek() {
                st.Pop()
            }

            if st.Empty() {
                arr = append(arr, asteroids[i])
            } else if st.Peek() == int(math.Abs(float64(asteroids[i]))) {
                st.Pop()
            }

          }
       } else {
            st.Push(asteroids[i])
       }
    }

    temp := NewStack[int]()
    for !st.Empty() {
        temp.Push(st.Pop())
    }

    for !temp.Empty() {
        arr = append(arr, temp.Pop())
    }

    return arr
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
    5 10 -5 -> 5, 10
    []

    10 2 -5 -> 10

    3 5 -6 2 -1 4 -> -6 2 4


    prev positive element -> right direction


    for every left moving elementt we find  previous greater element
    for every right movigin lement we find next smaller element

    3 5 -6 2 -1 4
*/