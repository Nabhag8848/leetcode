/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
        return root
    }

    queue := NewQueue[*Node]()
    queue.Enqueue(root)

    for !queue.Empty() {
        size := queue.Size()
        var node *Node = queue.Front()

        for size > 0 {
            node = queue.Dequeue()
            
            if node.Left != nil {
                queue.Enqueue(node.Left)
            }

            if node.Right != nil {
                queue.Enqueue(node.Right)
            }

            size--
            if size > 0 {
                node.Next = queue.Front()
            }
        }  
    }

    return root
}

type Queue[T any] struct {
	data  []T
	head  int
	count int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 4),
	}
}

func (q *Queue[T]) cap() int {
	return len(q.data)
}

func (q *Queue[T]) resize(newCap int) {
	newData := make([]T, newCap)
	for i := 0; i < q.count; i++ {
		newData[i] = q.data[(q.head+i)%q.cap()]
	}
	q.data = newData
	q.head = 0
}

func (q *Queue[T]) Enqueue(value T) {
	if q.count == q.cap() {
		q.resize(q.cap() * 2)
	}
	q.data[(q.head+q.count)%q.cap()] = value
	q.count++
}

func (q *Queue[T]) Dequeue() T {
	var zero T
	if q.Empty() {
		return zero
	}
	val := q.data[q.head]
	q.data[q.head] = zero
	q.head = (q.head + 1) % q.cap()
	q.count--
	return val
}

func (q *Queue[T]) Front() T {
	var zero T
	if q.Empty() {
		return zero
	}
	return q.data[q.head]
}

func (q *Queue[T]) Size() int {
	return q.count
}

func (q *Queue[T]) Empty() bool {
	return q.count == 0
}

