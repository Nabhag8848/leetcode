func orangesRotting(grid [][]int) int {
	arr := grid
	dq := NewDeque[Cell]()

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == 2 {
				dq.PushBack(Cell{row: i, col: j})
			}
		}
	}

	result := bfs(arr, dq)

	is_rotten := true
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == 1 {
				is_rotten = false
			}
		}
	}

	if is_rotten {
		return result
	}

	return -1
}


func bfs(grid [][]int, dq *Deque[Cell]) int {
	result := 0
	rows := len(grid)
	cols := len(grid[0])

	directions := []Cell{
		{row: -1, col: 0},
		{row: 1, col: 0},
		{row: 0, col: -1},
		{row: 0, col: 1},
	}

	for !dq.Empty() {
		levelSize := dq.Size()
		rotted := false

		for i := 0; i < levelSize; i++ {
			top := dq.PopFront()

			for _, direction := range directions {
				nextRow := top.row + direction.row
				nextCol := top.col + direction.col

				if nextRow < 0 || nextRow >= rows ||
					nextCol < 0 || nextCol >= cols ||
					grid[nextRow][nextCol] != 1 {
					continue
				}

				grid[nextRow][nextCol] = 2
				dq.PushBack(Cell{row: nextRow, col: nextCol})
				rotted = true
			}
		}

		if rotted {
			result++
		}
	}

	return result
}

type Cell struct {
    row int
    col int
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

