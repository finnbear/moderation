package radix

const (
	alphabet    = 26
	longestWord = 25
	chMax       = 1 + byte('z') - byte('a')
	chOffset    = byte('a')
)

type Tree struct {
	root   *Node
	length int
}

func New() Tree {
	return Tree{root: &Node{}}
}

func (tree *Tree) Root() *Node {
	return tree.root
}

func (tree *Tree) Add(word string, data uint32) {
	current := tree.root
	for i := 0; i < len(word); i++ {
		next := current.Next(word[i])
		if next == nil {
			next = &Node{depth: byte(i + 1), start: word[0]}
			current.children[word[i]-chOffset] = next
			current.hasChildren = true
		}
		current = next
	}
	current.data = data
	if !current.word {
		current.word = true
		tree.length++
	}
}

func (tree *Tree) get(word string) (node *Node) {
	current := tree.root
	for i := 0; i < len(word); i++ {
		current = current.Next(word[i])
		if current == nil {
			return
		}
	}
	return current
}

func (tree *Tree) Get(word string) (data uint32) {
	node := tree.get(word)
	if node != nil && node.word {
		data = node.data
	}
	return
}

func (tree *Tree) Contains(word string) (contains bool) {
	node := tree.get(word)
	return node != nil && node.word
}

// Preorder
func (tree *Tree) Traverse(callback func(string, uint32)) {
	tree.root.traverse(&[longestWord]byte{}, 0, callback)
}

func (tree *Tree) Len() (length int) {
	return tree.length
}
