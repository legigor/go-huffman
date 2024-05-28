package huffman

import (
	"github.com/legigor/go-huffman/pkg/huffman/bitbucket"
	"github.com/legigor/go-huffman/pkg/huffman/freqtable"
)

func compressAndDecompress(p []byte) ([]byte, error) {
	freq := freqtable.Initialize(p)
	root := createFreqTree(freq)

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

	freqTableCompressed := freq.Serialize()

	// TODO: return concatenated encoded + freqTable

	// Decompress

	// TODO: reconstruct freqTable and tree from encoded + freqTable

	freq = freqtable.Table{}
	freq.Deserialize(freqTableCompressed)

	root = createFreqTree(freq)

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
