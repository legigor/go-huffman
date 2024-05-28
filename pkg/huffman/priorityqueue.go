package huffman

import (
	"bytes"
	"container/heap"
	"encoding/gob"
	"fmt"
	"github.com/legigor/go-huffman/pkg/huffman/freqtable"
)

type huffmanNode struct {
	byte        byte
	freq        int
	left, right *huffmanNode
}

type priorityQueue []*huffmanNode

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
	*q = append(*q, x.(*huffmanNode))
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func (q *huffmanNode) serialize() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(q)
	if err != nil {
		return nil, fmt.Errorf("failed to encode huffmanNode: %w", err)
	}
	return buf.Bytes(), nil
}

func createFreqTree(freq freqtable.Table) *huffmanNode {
	pq := &priorityQueue{}
	heap.Init(pq)
	for b, freq := range freq {
		heap.Push(pq, &huffmanNode{byte: b, freq: freq})
	}

	for pq.Len() > 1 {
		left := heap.Pop(pq).(*huffmanNode)
		right := heap.Pop(pq).(*huffmanNode)
		heap.Push(pq, &huffmanNode{byte: 0, freq: left.freq + right.freq, left: left, right: right})
	}

	return heap.Pop(pq).(*huffmanNode)
}
