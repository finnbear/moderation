package radix

type Queue struct {
	// For correctness, at least longestWord * 2 because some characters turn into 2 matches
	// In reality, there are no profanities that are missed by reducing this to a lower level
	// For performance, should be a power of 2 so that modulo division is faster
	Storage    [32]Match

	// Length of abstract queue
	length     int

	// Indices of next read/write
	readIndex  int
	writeIndex int
}

const debug = false

// appends to back
func (queue *Queue) Append(match Match) {
	queue.length++

	if debug && queue.length > len(queue.Storage) {
		panic("queue too small")
	}

	queue.Storage[queue.writeIndex] = match
	queue.writeIndex = (queue.writeIndex + 1) % len(queue.Storage)
}

// appends to back if queue does not already contain node
func (queue *Queue) AppendUnique(match Match) {
	for i := 0; i < queue.length; i++ {
		idx := (queue.readIndex + i) % len(queue.Storage)
		if queue.Storage[idx].EqualsExceptLength(match) {
			// Not unique so return early
			return
		}
	}

	queue.Append(match)
}

// removes from front
func (queue *Queue) Remove() (match Match) {
	queue.length--

	if debug && queue.length < 0 {
		panic("queue out of range")
	}

	match = queue.Storage[queue.readIndex]
	queue.readIndex = (queue.readIndex + 1) % len(queue.Storage)
	return
}

func (queue *Queue) Clear() {
	queue.length = 0
	queue.readIndex = 0
	queue.writeIndex = 0
}

func (queue *Queue) Len() int {
	return queue.length
}
