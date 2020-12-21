package radix

type Buffer struct {
    Storage [longestWord * 2]*Node // * 2 because some characters turn into 2 matches
    index   int
}

func (buffer *Buffer) Append(node *Node) {
    buffer.Storage[buffer.index] = node
    buffer.index++
    if buffer.index >= len(buffer.Storage) {
        buffer.index = 0
    }
}

func (buffer *Buffer) Clear() {
    buffer.index = 0
}

func (buffer *Buffer) Get(index int) *Node {
    return buffer.Storage[index]
}

func (buffer *Buffer) Len() int {
    return buffer.index
}
