type Node struct {
    key int
    val int
    prev *Node
    next *Node
}

type LinkedList struct {
    head *Node
    tail *Node
}

type LRUCache struct {
    data map[int]*Node
    list *LinkedList
    capacity int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        data: make(map[int]*Node),
        list : NewLinkedList(),
        capacity: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.data[key]; ok {
        this.list.MoveToFront(node)
        return node.val
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
   
    if node, ok := this.data[key]; ok {
        this.list.DeleteNode(node) 
        newNode := this.list.InsertAtHead(key, value)
        this.data[key] = newNode
    } else {
        
        if this.capacity == 0 {
            delete(this.data, this.list.tail.key)
            this.list.DeleteNode(this.list.tail)
            newNode := this.list.InsertAtHead(key, value)
            this.data[key] = newNode
        } else {
            newNode := this.list.InsertAtHead(key, value)
            this.data[key] = newNode
            this.capacity--
        }
    }
}


func NewLinkedList() *LinkedList {
    return &LinkedList{
        head: nil,
        tail: nil,
    }
}

func (this *LinkedList) DeleteNode(node *Node) {
    if this.head == this.tail {
        this.head = nil
        this.tail = nil

        return 
    }

    if this.head == node {
        newHead := this.head.next
        this.head.next = nil
        newHead.prev = nil
        this.head = newHead

        return 
    }

    if this.tail == node {
        newTail := node.prev
        node.prev = nil
        newTail.next = nil
        this.tail = newTail

        return 
    }

    prevBefore := node.prev
    nextBefore := node.next

    node.prev = nil
    node.next = nil

    prevBefore.next = nextBefore
    nextBefore.prev = prevBefore
}

func (this *LinkedList) InsertAtHead(key int, val int) *Node {

    node := &Node{
        key: key,
        val: val,
        prev: nil,
        next: nil,
    }

    if this.head == nil {
        this.head = node 
        this.tail = node

        return this.head
    }


    node.next = this.head
    this.head.prev = node
    this.head = node

    return this.head
}

func (this *LinkedList) MoveToFront(node *Node) {
    if this.head == node {
        return 
    }

    if  this.tail == node {
        newTail := this.tail.prev
        newTail.next = nil
        this.tail.prev = nil

        newHead := this.tail
        this.tail = newTail


        newHead.next = this.head
        this.head.prev = newHead
        this.head = newHead 

        return   
    }

    prevBefore := node.prev
    nextBefore := node.next

    prevBefore.next = nextBefore
    nextBefore.prev = prevBefore

    node.prev = nil
    node.next = this.head
    this.head.prev = node
    this.head = node
}

func (this *LinkedList) InsertAtTail(key int, val int) {

    if this.tail == nil {
        this.InsertAtHead(key, val)
        return 
    }

    node := &Node{
        key: key,
        val: val,
        prev: nil,
        next: nil,
    }

    node.prev = this.tail
    this.tail.next = node
    this.tail = node
}



/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */