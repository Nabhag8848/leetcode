func nextGreaterElements(nums []int) []int {
    st := NewStack[int]()
    hash_map := make(map[int]int)
    ans := make([]int, len(nums))

    for i := len(nums) - 1; i >= 0; i-- {
        if st.Empty() {
            hash_map[i] = nums[i]
        } else if nums[i] < st.Peek() {
            ans[i] = st.Peek()
        } else {
            for !st.Empty() && nums[i] >= st.Peek() {
                st.Pop()
            }

            if st.Empty() {
               hash_map[i] = nums[i]
            } else {
                ans[i] = st.Peek()
            }
        }

        st.Push(nums[i])
    }

    for !st.Empty(){
        st.Pop()
    }

    for i := range nums {
        if value,ok := hash_map[i]; ok {
            is_break := false
            for j := 0; j < i; j++ {
                if nums[j] > value {
                    ans[i] = nums[j]
                    delete(hash_map, i)
                    is_break = true
                    break
                }
            }

            if !is_break {
                ans[i] = -1
            }
        }
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

/*
    1, 2, 1
    2 -1 -1

    stack<[]int>
    hash_map<index, (arr_value)>

    1,2,3,4,3
    2 3 4 -1 -1

    5 4 3 2 1
    -1 -1 -1 -1 -1

*/