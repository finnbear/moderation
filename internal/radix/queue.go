package radix

type Queue struct {
	Storage    [32]*Node // at least longestWord * 2 because some characters turn into 2 matches
	length     int
	readIndex  int
	writeIndex int
}

const debug = false

// appends to back
func (queue *Queue) Append(node *Node) {
	queue.length++

	if debug {
		if queue.length > len(queue.Storage) {
			panic("queue too small")
		}
	}

	queue.Storage[queue.writeIndex] = node
	queue.writeIndex = (queue.writeIndex + 1) % len(queue.Storage)
}

// appends to back if queue does not already contain node
func (queue *Queue) AppendUnique(node *Node) {
	unique := true
	for i := 0; i < queue.length; i++ {
		idx := (queue.readIndex + i) % len(queue.Storage)
		if queue.Storage[idx] == node {
			unique = false
			break
		}
	}

	if !unique {
		return
	}

	queue.Append(node)
}

// removes from front
func (queue *Queue) Remove() (node *Node) {
	queue.length--

	if debug {
		if queue.length < 0 {
			panic("queue out of range")
		}
	}

	node = queue.Storage[queue.readIndex]
	queue.readIndex = (queue.readIndex + 1) % len(queue.Storage)
	return
}

func (queue *Queue) Clear() {
	queue.length = 0
}

func (queue *Queue) Len() int {
	return queue.length
}
