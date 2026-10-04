func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand) % groupSize != 0 {
        return false
    }

    pq := NewOrderedPriorityQueue[int](hand)
    dq := NewDeque[int]()

    for !pq.Empty() || !dq.Empty() {
        for !dq.Empty() {
            pq.Push(dq.PopBack())
        }

        prev, _ := pq.Pop()
        size := 1
        for size < groupSize {
            top, exist := pq.Pop()

            if !exist {
                return false
            }

            if top == prev + 1 {
                prev = top
                size++
            } else if top == prev {
                dq.PushBack(top)
            } else {
                return false
            }
        }
    }

    return true
}

type Comparator[T any] func(a, b T) int

type Heap[T any] struct {
	data       []T
	comparator Comparator[T]
}


func NewHeap[T any](data []T, comparator Comparator[T]) *Heap[T] {
	return &Heap[T]{
		data:       data,
		comparator: comparator,
	}
}

func (h *Heap[T]) heapify() {
    length := len(h.data)

    for i := length / 2 - 1; i >= 0; i-- {
        h.downHeap(i)
    }
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) Empty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Empty() {
		var zero T
		return zero, false
	}

	return h.data[0], true
}

func (h *Heap[T]) parent(index int) (int, bool) {
	if index <= 0 || index >= h.Size() {
		return -1, false
	}

	return (index - 1) / 2, true
}

func (h *Heap[T]) leftChild(index int) (int, bool) {
	child := 2*index + 1

	if index < 0 || child >= h.Size() {
		return -1, false
	}

	return child, true
}

func (h *Heap[T]) rightChild(index int) (int, bool) {
	child := 2*index + 2

	if index < 0 || child >= h.Size() {
		return -1, false
	}

	return child, true
}

func (h *Heap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.upHeap(h.Size() - 1)
}

func (h *Heap[T]) upHeap(index int) {
	parent, exists := h.parent(index)

	if !exists ||
		h.comparator(h.data[index], h.data[parent]) >= 0 {
		return
	}

	h.data[index], h.data[parent] =
		h.data[parent], h.data[index]

	h.upHeap(parent)
}

func (h *Heap[T]) Pop() (T, bool) {
	if h.Empty() {
		var zero T
		return zero, false
	}

	result := h.data[0]
	last := h.Size() - 1

	h.data[0] = h.data[last]
	h.data = h.data[:last]

	if !h.Empty() {
		h.downHeap(0)
	}

	return result, true
}

func (h *Heap[T]) downHeap(index int) {
	best := index

	left, hasLeft := h.leftChild(index)
	if hasLeft &&
		h.comparator(h.data[left], h.data[best]) < 0 {
		best = left
	}

	right, hasRight := h.rightChild(index)
	if hasRight &&
		h.comparator(h.data[right], h.data[best]) < 0 {
		best = right
	}

	if best == index {
		return
	}

	h.data[index], h.data[best] =
		h.data[best], h.data[index]

	h.downHeap(best)
}


type PriorityQueue[T any] struct {
	heap *Heap[T]
}

func NewPriorityQueue[T any](
    data []T,
	comparator Comparator[T],
) *PriorityQueue[T] {
    heap := NewHeap(data, comparator)
    heap.heapify()

	return &PriorityQueue[T]{
		heap: heap,
	}
}

func NewOrderedPriorityQueue[T cmp.Ordered](data []T) *PriorityQueue[T] {
	return NewPriorityQueue(data, func(a, b T) int {
		return cmp.Compare(a, b)
	})
}

func (pq *PriorityQueue[T]) Push(value T) {
	pq.heap.Push(value)
}

func (pq *PriorityQueue[T]) Pop() (T, bool) {
	return pq.heap.Pop()
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	return pq.heap.Peek()
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.heap.Size()
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.heap.Empty()
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

/*
    1,2,3,6,2,3,4,7,8

    1 2 3

    2 
*/