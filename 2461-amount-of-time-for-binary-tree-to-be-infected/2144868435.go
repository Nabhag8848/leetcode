/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
    target := findTarget(root, start)
    parent := findNodeToParent(root)
   
    visited := make(map[*TreeNode]struct{})
    var distance int = -1
    visited[target] = struct{}{}
    queue := NewQueue[*TreeNode]()
    queue.Enqueue(target)

    for !queue.Empty() {
        size := queue.Size()

        distance++
        for size > 0 {
            top := queue.Dequeue()
            _,left_visited := visited[top.Left]
            if top.Left != nil && !left_visited {
                queue.Enqueue(top.Left)
                visited[top.Left] = struct{}{}
            }
            _, right_visited := visited[top.Right]
            if top.Right != nil && !right_visited {
                queue.Enqueue(top.Right)
                visited[top.Right] = struct{}{}
            } 

            parent_node, is_exist := parent[top] 
            _, parent_visited := visited[parent_node]
            if is_exist && !parent_visited {
                queue.Enqueue(parent_node)
                visited[parent_node] = struct{}{}
            }

            size--
        }
    }

    return distance
}

func findNodeToParent(root *TreeNode) map[*TreeNode]*TreeNode {
    parent := make(map[*TreeNode]*TreeNode)
    queue := NewQueue[*TreeNode]()
    queue.Enqueue(root)

    for !queue.Empty() {
        size := queue.Size()

        for size > 0 {
            top := queue.Dequeue()

            if top.Left != nil {
                parent[top.Left] = top
                queue.Enqueue(top.Left)
            }

            if top.Right != nil {
                parent[top.Right] = top
                queue.Enqueue(top.Right)
            }

            size--
        }
    }

    return parent
}

func findTarget(root *TreeNode, start int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val == start {
        return root
    }

    left := findTarget(root.Left, start)
    if left != nil {
        return left
    }

    right := findTarget(root.Right, start)
    if right != nil {
        return right
    }

    return nil
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