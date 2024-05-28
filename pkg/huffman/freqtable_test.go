package huffman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serialize_deserialize(t *testing.T) {
	freq := freqTable{'a': 1, 'b': 2, 'c': 3}

	serialized := freq.serialize()

	deserialized := freqTable{}
	deserialized.deserialize(serialized)

	assert.Equal(t, freq, deserialized)
}
