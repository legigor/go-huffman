package huffman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serialize_and_deserialize(t *testing.T) {

	freq := []byte("the frequency of the letters in this sentence")
	encoded := []byte("the encoded data of the sentence")

	archive := newArchivePage(freq, encoded)
	archiveData := archive.serialize()

	decodedArchive := deserialize(archiveData)

	assert.Equal(t, archive.codeLength, decodedArchive.codeLength)
	assert.Equal(t, archive.encodedSegmentLength, decodedArchive.encodedSegmentLength)
	assert.Equal(t, string(freq), string(decodedArchive.codeSegment))
	assert.Equal(t, string(encoded), string(decodedArchive.encodedSegment))
}
