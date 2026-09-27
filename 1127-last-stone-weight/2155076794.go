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

    current := h.end

    for {
        parent, exists := h.parent(current)
        if !exists || h.data[parent] >= h.data[current] {
            break
        }

        h.data[current], h.data[parent] =
            h.data[parent], h.data[current]

        current = parent
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

    for current := 0; current <= h.end; {
        largest := current

        left, hasLeft := h.leftChild(current)
        if hasLeft && h.data[left] > h.data[largest] {
            largest = left
        }

        right, hasRight := h.rightChild(current)
        if hasRight && h.data[right] > h.data[largest] {
            largest = right
        }

        if largest == current {
            break
        }

        h.data[current], h.data[largest] =
            h.data[largest], h.data[current]

        current = largest
    }

    return result, true
}