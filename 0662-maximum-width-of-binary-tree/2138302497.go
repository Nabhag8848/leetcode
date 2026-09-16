/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Pair struct {
    idx int
    node *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
    result := 0  
    if root == nil {
        return result
    }   

    queue := NewQueue[*Pair]()
    queue.Enqueue(&Pair{
        idx: 0,
        node:root,
    })

    for !queue.Empty() {
        size := queue.Size()
        var mini int
        var maxi int
        var min_idx int = queue.Front().idx
        for i:=0;i < size; i++ {
            top := queue.Dequeue()
            curr_idx := top.idx - min_idx
            node := top.node

            if i == 0 {
                mini = curr_idx
            }

            if i == size - 1 {
                maxi = curr_idx
            }

            if node.Left != nil {
                queue.Enqueue(&Pair{
                    node: node.Left,
                    idx: 2 * curr_idx + 1,
                })
            }

            if node.Right != nil {
                queue.Enqueue(&Pair{
                    node: node.Right,
                    idx: 2 * curr_idx + 2,
                })
            }
        }

        result = max(result, maxi - mini + 1)
    }

    return result
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