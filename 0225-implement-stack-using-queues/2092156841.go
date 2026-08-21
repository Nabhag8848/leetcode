type Queue struct {
	data []int
	start int
	end int
	size int
}

func NewQueue() *Queue {
	return &Queue{data: make([]int, 0), start: -1, end: -1, size: 0 }
}

func (q *Queue) Peek() int {
    if q.start < 0 {
	    return -1
    }

    return q.data[q.start]
}

func (q *Queue) Enqueue(value int)  {
	if q.size == 0 {
		q.start++
		q.end++
	} else {
        q.end++
    }

	q.data = append(q.data, value)
	q.size++
}

func (q *Queue) Dequeue() int { 
    if q.start < 0 {
        return -1
    }

	value := q.data[q.start]
	q.size--

	if q.size == 0 {
		q.start = -1
		q.end = -1
        q.data = make([]int, 0)
	} else {
        q.start++
    }

    return value
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

type MyStack struct {
    queue *Queue
}


func Constructor() MyStack {
    return MyStack{queue: NewQueue()}
}


func (this *MyStack) Push(x int)  {
    this.queue.Enqueue(x)

    from := 1
    to := this.queue.Size()

    for from < to {
        this.queue.Enqueue(this.queue.Peek())
        this.queue.Dequeue()
        from++
    }

}


func (this *MyStack) Pop() int {
    return this.queue.Dequeue()
}


func (this *MyStack) Top() int {
    return this.queue.Peek()
}


func (this *MyStack) Empty() bool {
    return this.queue.Empty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */