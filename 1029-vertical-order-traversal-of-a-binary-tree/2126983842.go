/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Node struct {
    ptr *TreeNode
    row int
    col int
}


func verticalTraversal(root *TreeNode) [][]int {
    hash_map := make(map[int]map[int][]int)
    queue := NewQueue[*Node]()
    max_level := -1
    queue.Enqueue(&Node{
        ptr: root,
        row: 0,
        col: 0,
    })
    for !queue.Empty() {
        size := queue.Size()
        max_level++

        for size > 0 {
            top := queue.Dequeue()
            node := top.ptr
            size--

            if node.Left != nil {
                queue.Enqueue(&Node{
                    ptr: node.Left,
                    row: top.row + 1, 
                    col: top.col - 1,
                })
            }

            if node.Right != nil {
                queue.Enqueue(&Node{
                    ptr: node.Right,
                    row: top.row + 1, 
                    col: top.col + 1,
                })
            }

            nodes, ok := hash_map[top.col]

            if ok {
                values, exist := nodes[top.row]

                if exist {
                    values = append(values, node.Val)
                    sort.Ints(values)
                    nodes[top.row] = values

                } else {
                    nodes[top.row] = []int{node.Val}
                }
            } else {
                hash_map[top.col] = make(map[int][]int)
                hash_map[top.col][top.row] = []int{node.Val}
            }
        }

    }

    result := make([][]int, 0)
    for col := -max_level; col <= max_level; col++ {
        arr := make([]int, 0)
        for row := 0; row <= max_level; row++ {
            if nodes, ok := hash_map[col][row]; ok {
                arr = append(arr, nodes...)
            }
        }
        if len(arr) > 0 {
            result = append(result, arr)
        }
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