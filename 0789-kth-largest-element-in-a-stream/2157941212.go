type KthLargest struct {
    heap *MinHeap
    kth int
}


func Constructor(k int, nums []int) KthLargest {
    heap := NewMinHeap()

    obj := KthLargest{
        heap: heap, 
        kth: k,
    }

    for i := range nums {
        obj.Add(nums[i])
    }

    return obj
}


func (this *KthLargest) Add(val int) int {
    this.heap.Insert(val)

    for this.heap.Size() > this.kth {
        this.heap.Delete()
    }

    return this.heap.Peek()
}

type MinHeap struct {
    data []int
}

func NewMinHeap() *MinHeap {
    return &MinHeap{
        data: []int{},
    }
}

func (this *MinHeap) downheap(parent, size int) {
    for {
        left := 2 * parent + 1
        right := 2 * parent + 2
        smallest := parent

        if left < size && this.data[left] < this.data[smallest] {
            smallest = left
        }

        if right < size && this.data[right] < this.data[smallest] {
            smallest = right
        }

        if smallest == parent {
            break
        }

        node_val := this.data[parent]
        this.data[parent] = this.data[smallest]
        this.data[smallest] = node_val
        parent = smallest
    }
}

func (this *MinHeap) Delete() int {
    length := this.Size()
    max := this.data[0]
    this.data[0] = this.data[length - 1]
    this.data = this.data[:length - 1]

    this.downheap(0, length - 1)
    return max
}


func (this *MinHeap) Peek() int{
    return this.data[0]
}

func (this *MinHeap) Empty() bool {
    return len(this.data) == 0
}

func (this *MinHeap) Size() int {
    return len(this.data)
}

func (this *MinHeap) Insert(val int) {
    this.data = append(this.data, val)
    this.upheap(this.Size() - 1)
}

func (this *MinHeap) parent(idx int) int {
    if idx > 0 && idx < this.Size() {
        return (idx - 1) / 2
    }

    return -1
}

func (this *MinHeap) upheap(child int) {
    for child > 0 {
        parent := this.parent(child)

        if this.data[parent] <= this.data[child] {
            break
        }

        this.data[parent], this.data[child] =
            this.data[child], this.data[parent]

        child = parent
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */