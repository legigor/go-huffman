package huffman

import (
	"container/heap"
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

	var encoded string
	for _, b := range p {
		encoded += codes[b]
	}

	// Decompress

	var decoded string
	node := root
	for _, bit := range encoded {
		if bit == '0' {
			node = node.left
		} else {
			node = node.right
		}
		if node.left == nil && node.right == nil {
			decoded += string(node.byte)
			node = root
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
