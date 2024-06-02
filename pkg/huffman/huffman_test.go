package huffman

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_compress_and_decompress(t *testing.T) {
	const testData = "beep boop beer!"

	t.Run("compress and decompress", func(t *testing.T) {
		compressed := compress([]byte(testData))
		decompressed := decompress(compressed)

		assert.Equal(t, testData, string(decompressed))
	})

	t.Run("compress and decompress with writer and reader", func(t *testing.T) {
		var compressed bytes.Buffer
		compressingWriter := NewWriter(&compressed)
		_, err := compressingWriter.Write([]byte(testData))
		require.NoError(t, err)
		err = compressingWriter.Close()
		require.NoError(t, err)

		decompressed := make([]byte, len(testData))
		decompressingReader := NewReader(&compressed)
		_, err = decompressingReader.Read(decompressed)
		require.NoError(t, err)
		err = decompressingReader.Close()
		require.NoError(t, err)

		assert.Equal(t, testData, string(decompressed))
	})
}
