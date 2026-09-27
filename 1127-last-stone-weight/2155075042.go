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
    if index < 0 || index > h.end {
        return -1, false
    }

    child := 2*index + 1
    if child > h.end {
        return -1, false
    }

    return child, true
}

func (h *MaxHeap) rightChild(index int) (int, bool) {
    if index < 0 || index > h.end {
        return -1, false
    }

    child := 2*index + 2
    if child > h.end {
        return -1, false
    }

    return child, true
}

func (h *MaxHeap) Push(value int) {
    h.data = append(h.data, value)
    h.end++

    idx := h.end

    for {
        parent, exists := h.parent(idx)
        if !exists || h.data[parent] >= h.data[idx] {
            break
        }

        h.data[parent], h.data[idx] = h.data[idx], h.data[parent]
        idx = parent
    }
}

func (h *MaxHeap) Pop() (int, bool) {
    if h.end == -1 {
        return 0, false
    }

    result := h.data[0]
    h.data[0] = h.data[h.end]
    h.data = h.data[:h.end]
    h.end--

    for idx := 0; idx <= h.end; {
        largest := idx

        left, hasLeft := h.leftChild(idx)
        if hasLeft && h.data[left] > h.data[largest] {
            largest = left
        }

        right, hasRight := h.rightChild(idx)
        if hasRight && h.data[right] > h.data[largest] {
            largest = right
        }

        if largest == idx {
            break
        }

        h.data[idx], h.data[largest] = h.data[largest], h.data[idx]
        idx = largest
    }

    return result, true
}