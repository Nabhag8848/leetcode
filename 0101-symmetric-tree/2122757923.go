/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    deque := NewDeque[*TreeNode]()

    if root.Left != nil {
        deque.PushBack(root.Left)
    }

    if root.Right != nil {
        deque.PushBack(root.Right)
    }

    for !deque.Empty() {
        front := deque.PopFront()
        back := deque.PopFront()
        
        if front == nil || back == nil {
            return front == back
        }

        if front.Val != back.Val {
            return false
        }

        if front.Left != nil && back.Right == nil {
            return false
        }

        if front.Left == nil && back.Right != nil {
            return false
        }

        if front.Right == nil && back.Left != nil {
            return false
        }

        if front.Right != nil && back.Left == nil {
            return false
        }

        if front.Left != nil {
            deque.PushBack(front.Left)
        }

        if back.Right != nil {
            deque.PushBack(back.Right)
        }

        if front.Right != nil {
            deque.PushBack(front.Right)  
        }

        if back.Left != nil {
            deque.PushBack(back.Left)
        }
    }

    return true
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