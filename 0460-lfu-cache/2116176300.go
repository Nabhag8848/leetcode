type LFUCache struct {
    freq int
    freq_list map[int]*LinkedList
    node map[int]*Node
    capacity int
}

type Node struct {
    key int
    val int
    freq int
    prev *Node
    next *Node
}

type LinkedList struct {
    head *Node
    tail *Node
}

func Constructor(capacity int) LFUCache {
   return LFUCache {
      freq: 0, 
      freq_list: make(map[int]*LinkedList, 0),
      node : make(map[int]*Node, 0),
      capacity: capacity,
   }
}


func (this *LFUCache) Get(key int) int {
    if node, ok := this.node[key]; ok {
        node_key := node.key
        node_val := node.val
        node_freq := node.freq

        this.freq_list[node_freq].DeleteNode(node)
        delete(this.node, node_key)

        if ll, is_exist := this.freq_list[node_freq]; is_exist {
                if ll.head == nil {
                    if node_freq == this.freq { 
                        this.freq++
                    }
                    delete(this.freq_list, node_freq)
                } 
        }


        if _, exist := this.freq_list[node_freq + 1]; !exist {
            this.freq_list[node_freq + 1] = NewLinkedList()
        }

        new_node := this.freq_list[node_freq + 1].InsertAtHead(node_key, node_val, node_freq + 1)
        this.node[key] = new_node

        return node_val
    }

    return -1
}

func (this *LFUCache) Put(key int, value int)  {
    if node, ok := this.node[key]; ok {
        node_key := node.key
        node_val := value
        node_freq := node.freq

        this.freq_list[node_freq].DeleteNode(node)
        delete(this.node, node_key)

        if ll, is_exist := this.freq_list[node_freq]; is_exist {
                if ll.head == nil {
                    if node_freq == this.freq {
                        this.freq++
                    }
                    delete(this.freq_list, node_freq)
                } 
        }


        if _, exist := this.freq_list[node_freq + 1]; !exist {
            this.freq_list[node_freq + 1] = NewLinkedList()
        }

        new_node := this.freq_list[node_freq + 1].InsertAtHead(node_key, node_val , node_freq + 1)
        this.node[key] = new_node

        return
    } 

    if this.capacity == 0 {
        ll, is_exist := this.freq_list[this.freq] 
        if is_exist {
            is_single_node := ll.head == ll.tail
            node_freq := ll.tail.freq
            node_key := ll.tail.key
            this.freq_list[node_freq].DeleteNode(ll.tail)
            delete(this.node, node_key)

            if is_single_node {
                delete(this.freq_list, node_freq)
            }

            this.capacity++
        }
        
    } 


    if _, exist := this.freq_list[1]; !exist {
        this.freq_list[1] = NewLinkedList()
    }


    new_node := this.freq_list[1].InsertAtHead(key, value, 1)
    this.node[key] = new_node
    this.freq = 1
    this.capacity--
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

func (this *LinkedList) InsertAtHead(key int, val int, freq int) *Node {

    node := &Node{
        key: key,
        val: val,
        freq: freq, 
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

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */