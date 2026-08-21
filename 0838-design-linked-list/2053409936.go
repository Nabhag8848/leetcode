type Node struct {
    value *int
    next *Node
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
    }

    if this.head == nil {
        this.head = &node
    } else {
        node.next = this.head
        this.head = &node
    }
}


func (this *MyLinkedList) AddAtTail(val int)  {
    fast := this.head

    node := Node{
        value: &val,
        next:nil,
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

	prev := this.head
	for i := 0; i < index-1; i++ {
		if prev == nil {
			return 
		}
		prev = prev.next
	}

	if prev == nil {
		return 
	}

	node := Node{value: &val, next: prev.next}
	prev.next = &node
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
		return
	}

	var prev *Node = this.head

	for i := 0; i < index-1; i++ {
		if prev == nil {
			return 
		}
		prev = prev.next
	}

	if prev == nil || prev.next == nil {
		return 
	}

	prev.next = prev.next.next
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