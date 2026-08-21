func subArrayRanges(nums []int) int64 {
    nse := findNextSmallerElement(nums)
    pse := findPrevSmallerElement(nums)
    nge := findNextGreaterElement(nums)
    pge := findPrevGreaterElement(nums)

    var total int64 = 0

    for i := range nums {
        leftMin := i - pse[i]
        rightMin := nse[i] - i
        leftMax := i - pge[i]
        rightMax := nge[i] - i

        maxi := rightMax * leftMax
        mini := leftMin * rightMin

        total += int64((maxi - mini) * nums[i])
    }

    return total
}

func findNextSmallerElement(nums []int) []int {
    result := make([]int, len(nums))
    st := NewStack[int]()
    
    for i := len(nums) - 1; i >= 0; i-- {
        for !st.Empty() && nums[st.Peek()] >= nums[i] {
            st.Pop()
        }

        if st.Empty() {
            result[i] = len(nums)
        } else {
            result[i] = st.Peek()
        }

        st.Push(i)
    }

    return result
}

func findPrevSmallerElement(nums []int) []int {
    result := make([]int, len(nums))
    st := NewStack[int]()
    
    for i := range nums {
        for !st.Empty() && nums[st.Peek()] > nums[i] {
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

func findNextGreaterElement(nums []int) []int {
    result := make([]int, len(nums))
    st := NewStack[int]()
    
    for i := len(nums) - 1; i >= 0; i-- {
        for !st.Empty() && nums[st.Peek()] <= nums[i] {
            st.Pop()
        }

        if st.Empty() {
            result[i] = len(nums)
        } else {
            result[i] = st.Peek()
        }

        st.Push(i)
    }

    return result
}

func findPrevGreaterElement(nums []int) []int {
    result := make([]int, len(nums))
    st := NewStack[int]()
    
    for i := range nums {
        for !st.Empty() && nums[st.Peek()] < nums[i] {
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
    1 2 3

    2 * 1 * 2 = 4

    ((2 - 1) * 2) * 2 = 4

    [1], range = largest - smallest = 1 - 1 = 0 
    [2], range = 2 - 2 = 0
    [3], range = 3 - 3 = 0
    [1,2], range = 2 - 1 = 1
    [2,3], range = 3 - 2 = 1
    [1,2,3], range = 3 - 1 = 2

    So the sum of all ranges is 0 + 0 + 0 + 1 + 1 + 2 = 4.

    // 2
    (3 - 1) * 1 - 0 = 2
    (2 - 1) * 2 = 2 

    // 1
    (3 - 0) * 1 = 3
    1 * (1) = 1. = -2

    // 3

    (3 - 2) * (2 - 1) = 1. == 2 * 3
    (3 - 2) * (2 + 1) = 3


*/