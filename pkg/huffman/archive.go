package huffman

import (
	"encoding/binary"
)

type archivePage struct {
	codeLength           int
	encodedSegmentLength int
	codeSegment          []byte
	encodedSegment       []byte
}

func newArchivePage(freq []byte, encoded []byte) *archivePage {
	return &archivePage{
		codeLength:           len(freq),
		encodedSegmentLength: len(encoded),
		codeSegment:          freq,
		encodedSegment:       encoded,
	}
}

func (a *archivePage) serialize() []byte {

	codeLength := intToBytes(a.codeLength)
	encodedLength := intToBytes(a.encodedSegmentLength)

	var arch []byte
	arch = append(arch, codeLength...)
	arch = append(arch, encodedLength...)
	arch = append(arch, a.codeSegment...)
	arch = append(arch, a.encodedSegment...)
	return arch
}

func deserialize(serialized []byte) *archivePage {

	codeLength := bytesToInt(serialized[0:4])
	encodedLength := bytesToInt(serialized[4:8])

	return &archivePage{
		codeLength:           codeLength,
		encodedSegmentLength: encodedLength,
		codeSegment:          serialized[8 : 8+codeLength],
		encodedSegment:       serialized[8+codeLength : 8+codeLength+encodedLength],
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
