func maxSlidingWindow(nums []int, k int) []int {
    result := make([]int, len(nums) - k + 1)
    deque := NewDeque[int]()

    for i := range nums {
        if !deque.Empty() && deque.Front() <= i - k {
            deque.PopFront()
        }

        for !deque.Empty() && nums[deque.Back()] <= nums[i] {
            deque.PopBack()
        }

        deque.PushBack(i)

        if i >= k - 1 {
            result[i - k + 1] = nums[deque.Front()]
        }
    }

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