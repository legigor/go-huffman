package bitbucket

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func asBinString(b byte) string {
	return fmt.Sprintf("%08b", b)
}

func Test_bits_bucket(t *testing.T) {

	t.Run("append bits as big-endian", func(t *testing.T) {
		bucket := Bucket{}
		bucket.Append(0, 1, 0, 1)

		require.Len(t, bucket.bits, 1)
		assert.Equal(t, asBinString(byte(0b_0101_0000)), asBinString(bucket.bits[0]))
	})

	t.Run("append bits one by one", func(t *testing.T) {
		bucket := Bucket{}
		bucket.Append(0)
		bucket.Append(0)
		bucket.Append(1)
		bucket.Append(0)
		bucket.Append(1)
		bucket.Append(1)

		require.Len(t, bucket.bits, 1)
		assert.Equal(t, asBinString(byte(0b_0010_1100)), asBinString(bucket.bits[0]))
	})

	seq := Bucket{}
	seq.Append(0, 0, 1, 1, 0, 1, 0, 1)
	seq.Append(1, 1, 1, 1)

	t.Run("append a new byte each 8 bits", func(t *testing.T) {
		require.Len(t, seq.bits, 2)
		assert.Equal(t, asBinString(byte(0b_0011_0101)), asBinString(seq.bits[0]))
		assert.Equal(t, asBinString(byte(0b_1111_0000)), asBinString(seq.bits[1]))
	})

	t.Run("read bits", func(t *testing.T) {
		expected := "0011010111110000"
		bits := seq.Iterator()
		total := 0
		var bitsRead string
		for bits.HasData() {
			if total > len(expected) {
				break
			}
			bitsRead = fmt.Sprintf("%s%d", bitsRead, bits.Read())
			total++
		}
		assert.Equal(t, expected, bitsRead)
	})
}
