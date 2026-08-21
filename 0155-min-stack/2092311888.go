type MinStack struct {
    data [][]int
    curr_min int
    top int
}


func Constructor() MinStack {
    return MinStack{ data: make([][]int, 0), top: -1, curr_min: math.MaxInt32 }
}


func (this *MinStack) Push(value int)  {
    if this.curr_min > value {
        this.curr_min = value
    }

    pair := []int{value, this.curr_min}
    this.data = append(this.data, pair)
    this.top++
}


func (this *MinStack) Pop() {
    this.top--

    if this.top < 0 {
        this.curr_min = math.MaxInt32
        this.data = make([][]int, 0)
    } else {
        this.data = this.data[:this.top + 1]
        this.curr_min = this.data[this.top][1]
    }
}


func (this *MinStack) Top() int {
    if this.top == -1 {
        return -1
    }

    return this.data[this.top][0]
}


func (this *MinStack) GetMin() int {
    return this.curr_min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */