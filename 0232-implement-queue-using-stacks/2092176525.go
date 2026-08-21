type Stack struct {
   data []int
   top int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0), top: -1 }
}

func (s *Stack) Push(value int) { 
	s.top++
	s.data = append(s.data, value)
 }

func (s *Stack) Pop() int {
    if s.top < 0 {
        return -1
    }
    
	value := s.data[s.top]
    s.top--

    if s.top == -1 {
        s.data = make([]int, 0)
    }
	return value
}

func (s *Stack) Peek() int {
	if s.top < 0 { 
		return -1
	}

	return s.data[s.top]
}

func (s *Stack) Empty() bool {
	return s.top == -1
}

func (s *Stack) Size() int {
	return s.top + 1
}

type MyQueue struct {
    st1 *Stack
    st2 *Stack
}


func Constructor() MyQueue {
    return MyQueue{ st1: NewStack(), st2: NewStack() }
}


func (this *MyQueue) Push(x int)  {
    this.st1.Push(x)
}


func (this *MyQueue) Pop() int {
    if !this.st2.Empty() {
        return this.st2.Pop()
    } else {
        for this.st1.Size() != 0 {
            this.st2.Push(this.st1.Pop())
        }
    }

    return this.st2.Pop()
}


func (this *MyQueue) Peek() int {
    if !this.st2.Empty() {
        return this.st2.Peek()
    } else {
        for this.st1.Size() != 0 {
            this.st2.Push(this.st1.Pop())
        }
    }

    return this.st2.Peek()
}


func (this *MyQueue) Empty() bool {
    length := this.st1.Size() + this.st2.Size()
    return length == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */