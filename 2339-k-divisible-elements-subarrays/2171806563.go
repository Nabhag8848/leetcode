func countDistinct(nums []int, k int, p int) int {
    trie := NewTrie(k, p)

    for i := range nums {
        trie.Insert(nums[i:])
    }

    return trie.result
}


type Trie struct {
    root *Node
    divisible int
    denominator int
    result int
}

type Node struct {
    links []*Node
}

func NewNode() *Node {
    return &Node{
        links: make([]*Node, 200),
    }
}

func (this *Node) containsKey(ch int) bool {
    return this.links[ch - 1] != nil
}

func (this *Node) putKey(ch int) {
    this.links[ch - 1] = NewNode()
}

func (this *Node) Next(ch int) *Node {
    return this.links[ch - 1]
}

func NewTrie(divisible, denominator int) *Trie {
    return &Trie{
        root: NewNode(),
        divisible: divisible,
        denominator: denominator,
        result: 0,
    }
}

func (this *Trie) Insert(nums []int) {
    node := this.root
    count := 0

    for _, num := range nums {
        if num % this.denominator == 0 {
            count++
        }

        if count > this.divisible {
            break
        }

        if !node.containsKey(num) {
            node.putKey(num)
            this.result++
        } 

        node = node.Next(num)
    }
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */