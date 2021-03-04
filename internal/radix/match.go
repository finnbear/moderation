package radix

type Match struct {
	Node     *Node
	Length   int  // how many matchable characters contributed
	Replaced bool // whether a replacement character contributed
	Separate bool // false if the match came after another caracter (no space/separation)
}

func (match Match) EqualsExceptLength(other Match) bool {
	return match.Node == other.Node && match.Replaced == other.Replaced && match.Separate == other.Separate
}
