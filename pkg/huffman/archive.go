package huffman

import (
	"encoding/binary"
)

type archivePage struct {
	codeLength           int
	encodedSegmentLength int
	originalLength       int
	codeSegment          []byte
	encodedSegment       []byte
}

func newArchivePage(code []byte, encoded []byte, originalLength int) *archivePage {
	return &archivePage{
		codeLength:           len(code),
		encodedSegmentLength: len(encoded),
		originalLength:       originalLength,
		codeSegment:          code,
		encodedSegment:       encoded,
	}
}

func (a *archivePage) serialize() []byte {

	codeLength := intToBytes(a.codeLength)
	encodedLength := intToBytes(a.encodedSegmentLength)
	originalLength := intToBytes(a.originalLength)

	var arch []byte
	arch = append(arch, codeLength...)
	arch = append(arch, encodedLength...)
	arch = append(arch, originalLength...)
	arch = append(arch, a.codeSegment...)
	arch = append(arch, a.encodedSegment...)
	return arch
}

func deserialize(serialized []byte) *archivePage {

	codeLength := bytesToInt(serialized[0:4])
	encodedLength := bytesToInt(serialized[4:8])
	originalLength := bytesToInt(serialized[8:12])

	return &archivePage{
		codeLength:           codeLength,
		encodedSegmentLength: encodedLength,
		originalLength:       originalLength,
		codeSegment:          serialized[12 : 12+codeLength],
		encodedSegment:       serialized[12+codeLength : 12+codeLength+encodedLength],
	}
}

func intToBytes(n int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	return buf
}

func bytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
}
