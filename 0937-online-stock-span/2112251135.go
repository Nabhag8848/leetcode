type StockSpanner struct {
    data []int
    s *Stack[int]
}


func Constructor() StockSpanner {
    return StockSpanner{data: make([]int, 0), s: NewStack[int]()}
}


func (this *StockSpanner) Next(price int) int {
    this.data = append(this.data, price)
    current_index := len(this.data) - 1

    for !this.s.Empty() && this.data[this.s.Peek()] <= price {
        this.s.Pop()
    }

    var ans int

    if this.s.Empty() {
        ans = current_index + 1
    }  else {
        ans = current_index - this.s.Peek()
    }

    this.s.Push(current_index)

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


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */