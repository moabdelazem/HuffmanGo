package huffman

type Node struct {
	Char  rune
	Freq  int
	Left  *Node
	Right *Node
}

func (n *Node) Less(other *Node) bool {
	return n.Freq < other.Freq
}
