package priorityqueue

import (
	"container/heap"
	"github.com/legigor/go-huffman/pkg/huffman/freqtable"
)

type HuffmanNode struct {
	freq int

	Byte        byte
	Left, Right *HuffmanNode
}

type priorityQueue []*HuffmanNode

func (q *priorityQueue) Len() int {
	return len(*q)
}

func (q *priorityQueue) Less(i, j int) bool {
	return (*q)[i].freq < (*q)[j].freq
}

func (q *priorityQueue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(*HuffmanNode))
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func CreateTree(freq freqtable.Table) *HuffmanNode {
	pq := &priorityQueue{}
	heap.Init(pq)
	for b, freq := range freq {
		heap.Push(pq, &HuffmanNode{Byte: b, freq: freq})
	}

	for pq.Len() > 1 {
		left := heap.Pop(pq).(*HuffmanNode)
		right := heap.Pop(pq).(*HuffmanNode)
		heap.Push(pq, &HuffmanNode{Byte: 0, freq: left.freq + right.freq, Left: left, Right: right})
	}

	return heap.Pop(pq).(*HuffmanNode)
}

func (q *HuffmanNode) Serialize() []byte {
	if q == nil {
		return nil
	}
	left := q.Left.Serialize()
	right := q.Right.Serialize()
	bytes := []byte{'{', q.Byte}
	bytes = append(bytes, left...)
	bytes = append(bytes, right...)
	bytes = append(bytes, '}')
	return bytes
}

func Deserialize(data []byte) *HuffmanNode {

	var d func(data []byte, pos int) (*HuffmanNode, int)
	d = func(data []byte, pos int) (*HuffmanNode, int) {
		pos++
		node := &HuffmanNode{Byte: data[pos]}
		pos++
		if data[pos] == '{' {
			node.Left, pos = d(data, pos)
			//pos++
		}
		if data[pos] == '{' {
			node.Right, pos = d(data, pos)
			//pos++
		}
		if data[pos] == '}' {
			pos++
		}
		return node, pos
	}

	root, _ := d(data, 0)
	return root
}
