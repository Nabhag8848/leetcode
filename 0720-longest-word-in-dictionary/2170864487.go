func longestWord(words []string) string {
    trie := NewTrie()

    for _, word := range words {
        trie.Insert(word)
    }

    return dfs(trie.root, []byte{}, "")
}

func dfs(node *Node, path []byte, best string) string {
    for i := 0; i < 26; i++ {
        child := node.links[i]

        if child == nil || child.end == 0 {
            continue
        }

        path = append(path, byte('a' + i))
        if len(path) > len(best) {
            best = string(path)
        }
        best = dfs(child, path, best)
        path = path[:len(path)-1]
    }

    return best
}

type Trie struct {
    root *Node
}

type Node struct {
    links []*Node
    prefix int
    end int
}

func NewNode() *Node {
    return &Node{
        links: make([]*Node, 26),
        prefix: 0, 
        end: 0,
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

func (this *Node) IncPrefix(ch rune) {
    node := this.links[ch - 'a']
    node.prefix++
}

func (this *Node) IncEnd() {
    this.end++
}

func NewTrie() *Trie {
    return &Trie{
        root: NewNode(),
    }
}

func (this *Trie) Insert(word string) {
    node := this.root

    for _, letter := range word {
        if !node.containsKey(letter) {
            node.putKey(letter)
        } else {
            node.IncPrefix(letter)
        }

        node = node.Next(letter)
    }

    node.IncEnd()
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */