package huffman

type PriorityQueue []*HeapNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Freq < pq[j].Freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	newItem := x.(*HeapNode)
	newItem.Index = len(*pq)
	*pq = append(*pq, newItem)
}

func (pq *PriorityQueue) Pop() interface{} {
	oldPQ := *pq
	n := len(oldPQ)
	item := oldPQ[n-1]
	item.Index = -1
	*pq = oldPQ[0 : n-1]
	return item
}
