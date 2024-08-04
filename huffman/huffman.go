package huffman

type Huffman struct {
	InternalChar []*Node
	Codes        map[byte]int
	FreqMap      map[int]int
}
