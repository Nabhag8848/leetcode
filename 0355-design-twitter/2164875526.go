type Tweet struct {
    tweetId   int
    timestamp int
    next      *Tweet
}

type Twitter struct {
    following map[int]map[int]bool
    tweets    map[int]*Tweet
    timestamp int
}

func Constructor() Twitter {
    return Twitter{
        following: make(map[int]map[int]bool),
        tweets:    make(map[int]*Tweet),
        timestamp: 0,
    }
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
    t.timestamp++

    tweet := &Tweet{
        tweetId:   tweetId,
        timestamp: t.timestamp,
        next:      t.tweets[userId],
    }

    // Insert at the head of the user's linked list
    t.tweets[userId] = tweet
}

func (t *Twitter) Follow(followerId int, followeeId int) {
    if followerId == followeeId {
        return
    }

    if t.following[followerId] == nil {
        t.following[followerId] = make(map[int]bool)
    }

    t.following[followerId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
    delete(t.following[followerId], followeeId)
}

func (t *Twitter) GetNewsFeed(userId int) []int {
    pq := NewPriorityQueue[*Tweet](func(a, b *Tweet) int {
        return cmp.Compare(b.timestamp, a.timestamp)
    })

    // Add user's most recent tweet
    if tweet := t.tweets[userId]; tweet != nil {
        pq.Push(tweet)
    }

    // Add most recent tweet from each followee
    for followeeId := range t.following[userId] {
        if followeeId == userId {
            continue
        }

        if tweet := t.tweets[followeeId]; tweet != nil {
            pq.Push(tweet)
        }
    }

    result := make([]int, 0, 10)

    for !pq.Empty() && len(result) < 10 {
        tweet, _ := pq.Pop()

        result = append(result, tweet.tweetId)

        // Push next tweet from the same user
        if tweet.next != nil {
            pq.Push(tweet.next)
        }
    }

    return result
}


type Comparator[T any] func(a, b T) int

type Heap[T any] struct {
	data       []T
	comparator Comparator[T]
}

func NewHeap[T any](comparator Comparator[T]) *Heap[T] {
	return &Heap[T]{
		data:       make([]T, 0),
		comparator: comparator,
	}
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) Empty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Empty() {
		var zero T
		return zero, false
	}

	return h.data[0], true
}

func (h *Heap[T]) parent(index int) (int, bool) {
	if index <= 0 || index >= h.Size() {
		return -1, false
	}

	return (index - 1) / 2, true
}

func (h *Heap[T]) leftChild(index int) (int, bool) {
	child := 2*index + 1

	if index < 0 || child >= h.Size() {
		return -1, false
	}

	return child, true
}

func (h *Heap[T]) rightChild(index int) (int, bool) {
	child := 2*index + 2

	if index < 0 || child >= h.Size() {
		return -1, false
	}

	return child, true
}

func (h *Heap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.upHeap(h.Size() - 1)
}

func (h *Heap[T]) upHeap(index int) {
	parent, exists := h.parent(index)

	if !exists ||
		h.comparator(h.data[index], h.data[parent]) >= 0 {
		return
	}

	h.data[index], h.data[parent] =
		h.data[parent], h.data[index]

	h.upHeap(parent)
}

func (h *Heap[T]) Pop() (T, bool) {
	if h.Empty() {
		var zero T
		return zero, false
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

func (h *Heap[T]) downHeap(index int) {
	best := index

	left, hasLeft := h.leftChild(index)
	if hasLeft &&
		h.comparator(h.data[left], h.data[best]) < 0 {
		best = left
	}

	right, hasRight := h.rightChild(index)
	if hasRight &&
		h.comparator(h.data[right], h.data[best]) < 0 {
		best = right
	}

	if best == index {
		return
	}

	h.data[index], h.data[best] =
		h.data[best], h.data[index]

	h.downHeap(best)
}


type PriorityQueue[T any] struct {
	heap *Heap[T]
}

func NewPriorityQueue[T any](
	comparator Comparator[T],
) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: NewHeap(comparator),
	}
}

func NewOrderedPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	return NewPriorityQueue(func(a, b T) int {
		return cmp.Compare(a, b)
	})
}

func (pq *PriorityQueue[T]) Push(value T) {
	pq.heap.Push(value)
}

func (pq *PriorityQueue[T]) Pop() (T, bool) {
	return pq.heap.Pop()
}

func (pq *PriorityQueue[T]) Peek() (T, bool) {
	return pq.heap.Peek()
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.heap.Size()
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.heap.Empty()
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */