package radix

type Node struct {
    children    [alphabet]*Node
    word        bool
    hasChildren bool
    data        int32
}

func (n *Node) Word() bool {
    return n.word
}

func (n *Node) Data() int32 {
    return n.data
}

func (node *Node) Next(next byte) *Node {
    return node.children[next - chOffset]
}

func (node *Node) traverse(word *[longestWord]byte, end int, callback func(string, int32)) {
    if node.word {
        callback(string(word[0:end]), node.data)
    }
    if node.hasChildren {
        for i, n := range node.children {
            if n != nil {
                word[end] = byte(i) + chOffset
                n.traverse(word, end + 1, callback)
            }
        }
    }
}
