type KthLargest struct {
    heap *MaxHeap
    kth int
}


func Constructor(k int, nums []int) KthLargest {
    heap := NewMaxHeap()

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
    this.heap.Insert(-val)

    for this.heap.Size() > this.kth {
        this.heap.Delete()
    }

    return -this.heap.Peek()
}

type MaxHeap struct {
    data []int
}

func NewMaxHeap() *MaxHeap {
    return &MaxHeap{
        data: []int{},
    }
}

func (this *MaxHeap) downheap(parent, size int) {
    for {
        left := 2 * parent + 1
        right := 2 * parent + 2
        largest := parent

        if left < size && this.data[left] > this.data[largest] {
            largest = left
        }

        if right < size && this.data[right] > this.data[largest] {
            largest = right
        }

        if largest == parent {
            break
        }

        node_val := this.data[parent]
        this.data[parent] = this.data[largest]
        this.data[largest] = node_val
        parent = largest
    }
}

func (this *MaxHeap) Delete() int {
    length := this.Size()
    max := this.data[0]
    this.data[0] = this.data[length - 1]
    this.data = this.data[:length - 1]

    this.downheap(0, length - 1)
    return max
}


func (this *MaxHeap) Peek() int{
    return this.data[0]
}

func (this *MaxHeap) Empty() bool {
    return len(this.data) == 0
}

func (this *MaxHeap) Size() int {
    return len(this.data)
}

func (this *MaxHeap) Insert(val int) {
    this.data = append(this.data, val)
    this.upheap(this.Size() - 1)
}

func (this *MaxHeap) parent(idx int) int {
    if idx > 0 && idx < this.Size() {
        return (idx - 1) / 2
    }

    return -1
}

func (this *MaxHeap) upheap(child int) {
    for child > 0 {
        parent := this.parent(child)

        if this.data[parent] >= this.data[child] {
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