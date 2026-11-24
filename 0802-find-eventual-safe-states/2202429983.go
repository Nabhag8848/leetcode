func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    revGraph := make([][]int, n)
    indegree := make([]int, n)

    for idx := range n {
        for _, edge := range graph[idx] {
            revGraph[edge] = append(revGraph[edge], idx)
            indegree[idx]++
        }
    }

    queue := NewDeque[int]()

    for idx := range n {
        if indegree[idx] == 0 {
            queue.PushBack(idx)
        }
    }

    var result []int

    for !queue.Empty() {
        top := queue.PopFront() 
        result = append(result, top)

        for _, node := range revGraph[top] {
            indegree[node]--
            if indegree[node] == 0 {
                queue.PushBack(node)
            }
        }
    }

    sort.Ints(result)
    return result
}

type Deque[T any] struct {
	data  []T
	head  int
	count int
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		data: make([]T, 4),
	}
}

func (d *Deque[T]) cap() int {
	return len(d.data)
}

func (d *Deque[T]) resize(newCap int) {
	newData := make([]T, newCap)
	for i := 0; i < d.count; i++ {
		newData[i] = d.data[(d.head+i)%d.cap()]
	}
	d.data = newData
	d.head = 0
}

func (d *Deque[T]) PushFront(value T) {
	if d.count == d.cap() {
		d.resize(d.cap() * 2)
	}
	d.head = (d.head - 1 + d.cap()) % d.cap()
	d.data[d.head] = value
	d.count++
}

func (d *Deque[T]) PushBack(value T) {
	if d.count == d.cap() {
		d.resize(d.cap() * 2)
	}
	d.data[(d.head+d.count)%d.cap()] = value
	d.count++
}

func (d *Deque[T]) PopFront() T {
	var zero T
	if d.Empty() {
		return zero
	}
	val := d.data[d.head]
	d.data[d.head] = zero
	d.head = (d.head + 1) % d.cap()
	d.count--
	return val
}

func (d *Deque[T]) PopBack() T {
	var zero T
	if d.Empty() {
		return zero
	}
	idx := (d.head + d.count - 1) % d.cap()
	val := d.data[idx]
	d.data[idx] = zero
	d.count--
	return val
}

func (d *Deque[T]) Front() T {
	var zero T
	if d.Empty() {
		return zero
	}
	return d.data[d.head]
}

func (d *Deque[T]) Back() T {
	var zero T
	if d.Empty() {
		return zero
	}
	return d.data[(d.head+d.count-1)%d.cap()]
}

func (d *Deque[T]) Size() int {
	return d.count
}

func (d *Deque[T]) Empty() bool {
	return d.count == 0
}
