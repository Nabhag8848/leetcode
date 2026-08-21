type Node struct {
    value *int
    next *Node
    prev *Node
}

type MyLinkedList struct {
    head *Node
}

func Constructor() MyLinkedList {
    return MyLinkedList{
        head: nil,
    }
}

func (this *MyLinkedList) Get(index int) int {
    first := this.head

    for i:=0; i < index; i++ {
        if first == nil {
            return -1
        }

        first = first.next
    }

	if first == nil {
		return -1
	}

    return *first.value
}

func (this *MyLinkedList) AddAtHead(val int)  {
    node := Node{
        value: &val,
        next:nil,
        prev:nil,
    }

    if this.head == nil {
        this.head = &node
    } else {
        node.next = this.head
        this.head.prev = &node
        this.head = &node
    }
}


func (this *MyLinkedList) AddAtTail(val int)  {
    fast := this.head

    node := Node{
        value: &val,
        next:nil,
        prev:nil,
    }

    if this.head == nil {
        this.head = &node
    } else {
        for fast != nil && fast.next != nil {
            fast = fast.next
        }

        fast.next = &node
    }
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index == 0 {
		this.AddAtHead(val)
		return
	}

	current := this.head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return 
		}
		current = current.next
	}

	if current == nil {
		return 
	}

	node := Node{value: &val, next: current.next, prev:current}
    
    if current.next != nil {
        current.next.prev = &node
    }

	current.next = &node
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 {
		return
	}

	if index == 0 {
		if this.head == nil {
			return
		}
		this.head = this.head.next

        if this.head != nil {
            this.head.prev = nil
        }
		return
	}

	var current *Node = this.head

	for i := 0; i < index-1; i++ {
		if current == nil {
			return 
		}
		current = current.next
	}

	if current == nil || current.next == nil {
		return 
	}

    next := current.next.next

    if next != nil {
        next.prev = current
    }
	current.next = current.next.next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */