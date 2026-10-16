type Pair struct {
    First  int
    Second int
}

type NestedPair struct {
    First  int
    Second Pair
}

func maximizeXor(nums []int, queries [][]int) []int {
    sort.Ints(nums)
    
    oQ := make([]NestedPair, 0)

    for idx, query := range queries {
        oQ = append(oQ, NestedPair{
            First: query[1],
            Second: Pair{
                First:  query[0],
                Second: idx,
            },
        })
    }

    sort.Slice(oQ, func (i, j int) bool {
        return oQ[i].First < oQ[j].First
    })

    result := make([]int, len(queries))
    trie := NewTrie()

    idx := 0
    length := len(nums)

    for _, query := range oQ {
        xi := query.Second.First
        ai := query.First

        for idx < length && nums[idx] <= ai {
            trie.Insert(nums[idx])
            idx++
        }

        if idx == 0 {
            result[query.Second.Second] = -1
        } else {
            result[query.Second.Second] = trie.getMax(xi)
        }
    }

    return result
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
