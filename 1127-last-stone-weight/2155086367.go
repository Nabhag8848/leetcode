func lastStoneWeight(stones []int) int {
    heap := NewMaxHeap()

    for _, stone := range stones {
        heap.Push(stone)
    }

    for heap.Size() > 1 {
        first, _ := heap.Pop()
        second, _ := heap.Pop()

        if first != second {
            heap.Push(first - second)
        }
    }

    result, _ := heap.Pop()
    return result
}

type MaxHeap struct {
    data []int
}

func NewMaxHeap() *MaxHeap {
    return &MaxHeap{
        data: make([]int, 0),
    }
}

func (h *MaxHeap) Size() int {
    return len(h.data)
}

func (h *MaxHeap) Empty() bool {
    return h.Size() == 0
}

func (h *MaxHeap) Peek() (int, bool) {
    if h.Empty() {
        return 0, false
    }

    return h.data[0], true
}

func (h *MaxHeap) parent(index int) (int, bool) {
    if index <= 0 || index >= h.Size() {
        return -1, false
    }

    return (index - 1) / 2, true
}

func (h *MaxHeap) leftChild(index int) (int, bool) {
    child := 2*index + 1
    if index < 0 || child >= h.Size() {
        return -1, false
    }

    return child, true
}

func (h *MaxHeap) rightChild(index int) (int, bool) {
    child := 2*index + 2
    if index < 0 || child >= h.Size() {
        return -1, false
    }

    return child, true
}

func (h *MaxHeap) Push(value int) {
    h.data = append(h.data, value)
    h.upHeap(h.Size() - 1)
}

func (h *MaxHeap) upHeap(index int) {
    parent, exists := h.parent(index)
    if !exists || h.data[parent] >= h.data[index] {
        return
    }

    h.data[index], h.data[parent] = h.data[parent], h.data[index]
    h.upHeap(parent)
}

func (h *MaxHeap) Pop() (int, bool) {
    if h.Empty() {
        return 0, false
    }

    result := h.data[0]
    last := h.Size() - 1

    h.data[0] = h.data[last]
    h.data = h.data[:last]

    if !h.Empty() {
        h.downHeap(0)
    }

    return result, true
}

func (h *MaxHeap) downHeap(index int) {
    largest := index

    left, hasLeft := h.leftChild(index)
    if hasLeft && h.data[left] > h.data[largest] {
        largest = left
    }

    right, hasRight := h.rightChild(index)
    if hasRight && h.data[right] > h.data[largest] {
        largest = right
    }

    if largest == index {
        return
    }

    h.data[index], h.data[largest] = h.data[largest], h.data[index]
    h.downHeap(largest)
}