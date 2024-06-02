package huffman

import (
	"github.com/legigor/go-huffman/pkg/huffman/bitbucket"
	"github.com/legigor/go-huffman/pkg/huffman/freqtable"
	"github.com/legigor/go-huffman/pkg/huffman/priorityqueue"
)

func compressAndDecompress(p []byte) ([]byte, error) {
	freq := freqtable.Initialize(p)
	root := priorityqueue.CreateTree(freq)

	codes := make(map[byte]string)
	var generateCodes func(node *priorityqueue.HuffmanNode, code string)
	generateCodes = func(node *priorityqueue.HuffmanNode, code string) {
		if node == nil {
			return
		}
		if node.Byte != 0 {
			codes[node.Byte] = code
			return
		}
		generateCodes(node.Left, code+"0")
		generateCodes(node.Right, code+"1")
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

	//freqTableCompressed := freq.Serialize()
	treeCompressed := root.Serialize()

	// TODO: return concatenated encoded + freqTable

	// Decompress

	// TODO: reconstruct freqTable and tree from encoded + freqTable

	root = priorityqueue.Deserialize(treeCompressed).Left

	var decoded string
	currentNode := root

	bitsIterator := encoded.Iterator()
	for bitsIterator.HasData() {
		bit := bitsIterator.Read()
		if bit == 0 {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
		if currentNode.Left == nil && currentNode.Right == nil {
			decoded += string(currentNode.Byte)
			currentNode = root
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
