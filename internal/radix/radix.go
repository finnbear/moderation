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

func New() *Tree {
	return NewFromList("")
}

func NewFromList(words string) (tree *Tree) {
	var (
		scanBuffer [longestWord]byte
		nodeCache  [longestWord + 1]*Node
	)

	tree = &Tree{root: &Node{}}

	nodeCache[0] = tree.root
	i := 0               // Index of current word
	previousLength := -2 // Length of previous word
	indexed := 0         // How many words were indexed
	p := tree.root

	for c := 0; c < len(words); c++ {
		ch := words[c] - chOffset
		if ch < chMax && i < longestWord {
			if scanBuffer[i] == ch {
				if i < previousLength {
					i++
					continue
				}
			} else {
				scanBuffer[i] = ch
			}

			i++

			if previousLength != -1 {
				previousLength = -1
				p = nodeCache[i-1]
			}

			q := &Node{}

			p.children[ch] = q
			nodeCache[i] = q
			p.hasChildren = true

			p = q
		} else if i > 0 {
			tree.length++
			p.word = true

			indexed++
			previousLength = i
			i = 0
		}
	}
	return
}

func (tree *Tree) Root() *Node {
	return tree.root
}

func (tree *Tree) Add(word string, data int32) {
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

func (tree *Tree) Remove(word string) {
	current := tree.root
	for i := 0; i < len(word); i++ {
		next := current.Next(word[i])
		if next == nil {
			break
		}
		if i == len(word)-1 {
			if next.word {
				tree.length--
			}
			if next.hasChildren {
				next.data = 0
				next.word = false
			} else {
				current.children[word[i]-chOffset] = nil
				// TODO: update hasChildren
				// TODO: update higher nodes
			}
			break
		}
		current = next
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

func (tree *Tree) Get(word string) (data int32) {
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
func (tree *Tree) Traverse(callback func(string, int32)) {
	tree.root.traverse(&[longestWord]byte{}, 0, callback)
}

func (tree *Tree) Len() (length int) {
	return tree.length
}
