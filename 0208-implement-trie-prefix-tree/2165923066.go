type Trie struct {
   root *Node
}

type Node struct {
    links []*Node
    flag bool
}

func NewNode() *Node{
    return &Node{
        links: make([]*Node, 26),
        flag: false,
    }
}

func (this *Node) containsKey(ch rune) bool {
    return this.links[ch - 'a'] != nil
}

func (this *Node) putKey(ch rune) {
    this.links[ch - 'a'] = NewNode()
}

func (this *Node) Next(ch rune) *Node {
    return this.links[ch - 'a']
}

func (this *Node) End() {
    this.flag = true
}

func (this *Node) IsEnd() bool {
    return this.flag
}

func Constructor() Trie {
    return Trie{
        root: NewNode(),
    }
}


func (this *Trie) Insert(word string)  {
    node := this.root

    for _, letter := range word {
        if !node.containsKey(letter) {
            node.putKey(letter)
        }

        node = node.Next(letter)
    }

    node.End()
}


func (this *Trie) Search(word string) bool {
    node := this.root

    for _, letter := range word {
        if !node.containsKey(letter) {
            return false
        }

        node = node.Next(letter)
    }

    return node.IsEnd()
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this.root

    for _, letter := range prefix {
        if !node.containsKey(letter) {
            return false
        }

        node = node.Next(letter)
    }

    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */