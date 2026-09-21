/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    queue *Queue[*TreeNode] 
}

func Constructor() Codec {
    return Codec{
        queue: NewQueue[*TreeNode](),
    }
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    str := ""

    if root == nil {
        return str
    }

    this.queue.Enqueue(root)

    for !this.queue.Empty() {
        size := this.queue.Size()
        for size > 0 {
            node := this.queue.Dequeue()
            if node == root {
                str = str + strconv.Itoa(node.Val)
            }

            size--
            
            if node.Left != nil {
                str = str + "," + strconv.Itoa(node.Left.Val)
                this.queue.Enqueue(node.Left)
            } else {
                str = str + ",#"
            }

            if node.Right != nil {
                str = str + "," + strconv.Itoa(node.Right.Val)
                this.queue.Enqueue(node.Right)
            } else {
                str = str + ",#"
            }
        }
    }

    return str
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	values := strings.Split(data, ",")
	val, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: val}

	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for i := 1; !queue.Empty(); i += 2 {
		node := queue.Dequeue()

		if values[i] != "#" {
			val, _ := strconv.Atoi(values[i])
			node.Left = &TreeNode{Val: val}
			queue.Enqueue(node.Left)
		}

		if values[i+1] != "#" {
			val, _ := strconv.Atoi(values[i+1])
			node.Right = &TreeNode{Val: val}
			queue.Enqueue(node.Right)
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

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