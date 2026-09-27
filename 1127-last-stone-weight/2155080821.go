func lastStoneWeight(stones []int) int {
    heap := NewMaxHeap()

    for _, stone := range stones {
        heap.Push(stone)
    }

    for {
        first, hasFirst := heap.Pop()
        if !hasFirst {
            return 0
        }

        second, hasSecond := heap.Pop()
        if !hasSecond {
            return first
        }

        if first != second {
            heap.Push(first - second)
        }
    }
}

type MaxHeap struct {
    data []int
    end  int
}

func NewMaxHeap() *MaxHeap {
    return &MaxHeap{
        data: make([]int, 0),
        end:  -1,
    }
}

func (h *MaxHeap) parent(index int) (int, bool) {
    if index <= 0 || index > h.end {
        return -1, false
    }

    return (index - 1) / 2, true
}

func (h *MaxHeap) leftChild(index int) (int, bool) {
    child := 2*index + 1
    if index < 0 || child > h.end {
        return -1, false
    }

    return child, true
}

func (h *MaxHeap) rightChild(index int) (int, bool) {
    child := 2*index + 2
    if index < 0 || child > h.end {
        return -1, false
    }

    return child, true
}

func (h *MaxHeap) Push(value int) {
    h.data = append(h.data, value)
    h.end++

    h.upHeap(h.end)
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
    if h.end == -1 {
        return 0, false
    }

    result := h.data[0]
    h.data[0] = h.data[h.end]
    h.data = h.data[:h.end]
    h.end--

    if h.end >= 0 {
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