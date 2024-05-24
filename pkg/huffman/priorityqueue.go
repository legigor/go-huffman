package huffman

import (
	"bytes"
	"encoding/gob"
	"fmt"
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
