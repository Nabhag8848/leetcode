func findKthLargest(nums []int, k int) int {
    heap := NewMaxHeap(nums)
    heap.convert()

    for k > 1 {
        heap.Delete()
        k--
    }

    return heap.Peek()
}

type MaxHeap struct {
    data []int
}

func NewMaxHeap(data []int) *MaxHeap {
    return &MaxHeap{
        data: data,
    }
}

func (this *MaxHeap) convert() {
    length := len(this.data)

    for i := length / 2 - 1; i >= 0; i-- {
        this.downheap(i, length)
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

func (this *MaxHeap) Peek() int{
    return this.data[0]
}

func (this *MaxHeap) Delete() int {
    length := len(this.data)
    max := this.data[0]
    this.data[0] = this.data[length - 1]
    this.data = this.data[:length - 1]

    this.downheap(0, length - 1)
    return max
}


