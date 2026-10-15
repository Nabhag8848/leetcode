func findMaximumXOR(nums []int) int {
    maxi := 0
    trie := NewTrie()

    for _, num := range nums {
        trie.Insert(num)
    }

    for _, num := range nums {
        maxi = max(maxi, trie.getMax(num))
    }

    return maxi
}

type Trie struct {
   root *Node
}

type Node struct {
    links []*Node
}

func NewNode() *Node{
    return &Node{
        links: make([]*Node, 2),
    }
}

func (this *Node) containsKey(bit int) bool {
    return this.links[bit] != nil
}

func (this *Node) putKey(bit int) {
    this.links[bit] = NewNode()
}

func (this *Node) Next(bit int) *Node {
    return this.links[bit]
}

func NewTrie() Trie {
    return Trie{
        root: NewNode(),
    }
}

func (this *Trie) Insert(num int)  {
    node := this.root

    for i := 31; i >= 0; i-- {
        bit := (num >> i) & 1

        if !node.containsKey(bit) {
            node.putKey(bit)
        }

        node = node.Next(bit)
    }
}

func (this *Trie) getMax(num int) int {
    node := this.root
    maxi := 0 

    for i := 31; i >= 0; i-- {
        bit := (num >> i) & 1

        if node.containsKey(1 - bit) {
            maxi = maxi | (1 << i)
            node = node.Next(1 - bit) 
        } else {
            node = node.Next(bit)
        }   
    }

    return maxi
}
