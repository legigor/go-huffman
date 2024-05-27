package huffman

import (
	"container/heap"
	"github.com/legigor/go-huffman/pkg/huffman/bitbucket"
)

func compressAndDecompress(p []byte) ([]byte, error) {
	freq := make(map[byte]int)
	for _, b := range p {
		freq[b]++
	}

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

	root := heap.Pop(pq).(*huffmanNode)

	codes := make(map[byte]string)
	var generateCodes func(node *huffmanNode, code string)
	generateCodes = func(node *huffmanNode, code string) {
		if node == nil {
			return
		}
		if node.byte != 0 {
			codes[node.byte] = code
			return
		}
		generateCodes(node.left, code+"0")
		generateCodes(node.right, code+"1")
	}
	generateCodes(root, "")

	encoded := bitbucket.Bucket{}
	enc := ""
	for _, b := range p {
		code := codes[b]
		for _, c := range code {
			if c == '0' {
				encoded.Append(0)
				enc += "0"
			} else {
				encoded.Append(1)
				enc += "1"
			}
		}
	}

	// Decompress

	var decoded string
	node := root

	bitsIterator := encoded.Iterator()
	for bitsIterator.HasData() {
		bit := bitsIterator.Read()
		if bit == 0 {
			node = node.left
		} else {
			node = node.right
		}
		if node.left == nil && node.right == nil {
			decoded += string(node.byte)
			node = root
		}

		// Must not process more bits than encoded
		if len(decoded) == len(p) {
			break
		}
	}

	return []byte(decoded), nil
}

func compress(p []byte) ([]byte, error) {
	return p, nil
}

func decompress(p []byte) []byte {
	return p
}
