package radix

type Queue struct {
	Storage    [longestWord * 2]*Node // * 2 because some characters turn into 2 matches
	length     int
	readIndex  int
	writeIndex int
}

// appends to back
func (queue *Queue) Append(node *Node) {
	queue.length++

	queue.Storage[queue.writeIndex] = node
	queue.writeIndex++
	if queue.writeIndex >= len(queue.Storage) {
		queue.writeIndex = 0
	}
}

// removes from front
func (queue *Queue) Remove() (node *Node) {
	queue.length--

	/*
	   if queue.length < 0 {
	       panic("queue out of range")
	   }
	*/

	node = queue.Storage[queue.readIndex]
	queue.readIndex++
	if queue.readIndex >= len(queue.Storage) {
		queue.readIndex = 0
	}
	return
}

func (queue *Queue) Clear() {
	queue.length = 0
}

func (queue *Queue) Len() int {
	return queue.length
}
