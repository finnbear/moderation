package radix

type Node struct {
	children    [alphabet]*Node
	word        bool
	hasChildren bool
	start       byte // starting character (not offset)
	depth       byte
	data        uint32
}

func (node *Node) Word() bool {
	return node.word
}

func (node *Node) Data() uint32 {
	return node.data
}

func (node *Node) Depth() int {
	return int(node.depth)
}

func (node *Node) Next(next byte) *Node {
	return node.children[next-chOffset]
}

func (node *Node) Start() byte {
	return node.start
}

func (node *Node) traverse(word *[longestWord]byte, end int, callback func(string, uint32)) {
	if node.word {
		callback(string(word[0:end]), node.data)
	}
	if node.hasChildren {
		for i, n := range node.children {
			if n != nil {
				word[end] = byte(i) + chOffset
				n.traverse(word, end+1, callback)
			}
		}
	}
}
