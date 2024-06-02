package priorityqueue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_serialization(t *testing.T) {

	t.Run("single node", func(t *testing.T) {
		n := &HuffmanNode{}
		s := n.Serialize()
		assert.Equal(t, []byte{'{', 0, '}'}, s)
	})

	t.Run("single node", func(t *testing.T) {
		n := &HuffmanNode{}
		s := n.Serialize()
		assert.Equal(t, []byte{'{', 0, '}'}, s)
	})

	t.Run("tree", func(t *testing.T) {
		n := &HuffmanNode{
			Byte: 0,
			Left: &HuffmanNode{
				Byte: 0,
				Left: &HuffmanNode{
					Byte: 'a',
				},
				Right: &HuffmanNode{
					Byte: 0,
					Left: &HuffmanNode{
						Byte: 'b',
					},
				},
			},
		}

		s := n.Serialize()
		assert.Equal(t, string([]byte{'{', 0, '{', 0, '{', 'a', '}', '{', 0, '{', 'b', '}', '}', '}', '}'}), string(s))
	})

	t.Run("advanced tree", func(t *testing.T) {
		n := &HuffmanNode{
			Byte: 0,
			Left: &HuffmanNode{
				Byte: 0,
				Left: &HuffmanNode{
					Byte: 98,
				},
				Right: &HuffmanNode{
					Byte: 0,
					Left: &HuffmanNode{
						Byte: 97,
					},
					Right: &HuffmanNode{
						Byte: 0,
						Left: &HuffmanNode{
							Byte: 33,
						},
						Right: &HuffmanNode{
							Byte: 114,
						},
					},
				},
			},
			Right: &HuffmanNode{
				Byte: 0,
				Left: &HuffmanNode{
					Byte: 101,
				},
				Right: &HuffmanNode{
					Byte: 0,
					Left: &HuffmanNode{
						Byte: 111,
					},
					Right: &HuffmanNode{
						Byte: 112,
					},
				},
			},
		}
		s := n.Serialize()
		expected := []byte{'{', 0, '{', 0, '{', 98, '}', '{', 0, '{', 97, '}', '{', 0, '{', 33, '}', '{', 114, '}', '}', '}', '}', '{', 0, '{', 101, '}', '{', 0, '{', 111, '}', '{', 112, '}', '}', '}', '}'}
		assert.Equal(t, string(expected), string(s))
	})
}

func Test_deserialization(t *testing.T) {

	t.Run("single node", func(t *testing.T) {
		testData := []byte{'{', 0, '}'}
		tree := Deserialize(testData)
		assert.Equal(t, byte(0), tree.Byte)
	})

	t.Run("single node - one child", func(t *testing.T) {
		testData := []byte{'{', 0, '{', 'a', '}', '}'}
		tree := Deserialize(testData)
		assert.Equal(t, byte(0), tree.Byte)
		assert.Equal(t, byte('a'), tree.Left.Byte)
		assert.Nil(t, tree.Right)
	})

	t.Run("single node - two child", func(t *testing.T) {
		testData := []byte{'{', 0, '{', 'a', '}', '{', 'b', '}', '}'}
		tree := Deserialize(testData)
		assert.Equal(t, byte(0), tree.Byte)
		assert.Equal(t, byte('a'), tree.Left.Byte)
		assert.Equal(t, byte('b'), tree.Right.Byte)
	})

	t.Run("tree", func(t *testing.T) {
		testData := []byte{'{', 0, '{', 'a', '}', '{', 0, '{', 'b', '}', '{', 'c', '}', '}', '}'}
		tree := Deserialize(testData)

		require.NotNil(t, tree.Left)
		require.NotNil(t, tree.Right)

		assert.Equal(t, byte(0), tree.Byte)
		assert.Equal(t, byte('a'), tree.Left.Byte)
		assert.Equal(t, byte(0), tree.Right.Byte)
		assert.Equal(t, byte('b'), tree.Right.Left.Byte)
		assert.Equal(t, byte('c'), tree.Right.Right.Byte)
	})

	t.Run("final boss tree", func(t *testing.T) {
		serialized := []byte{'{', 0, '{', 0, '{', 98, '}', '{', 0, '{', 97, '}', '{', 0, '{', 33, '}', '{', 114, '}', '}', '}', '}', '{', 0, '{', 101, '}', '{', 0, '{', 111, '}', '{', 112, '}', '}', '}', '}'}
		tree := Deserialize(serialized)

		assert.Equal(t, byte(0), tree.Byte)
		assert.Equal(t, byte(0), tree.Left.Byte)
		assert.Equal(t, byte(98), tree.Left.Left.Byte)
		assert.Equal(t, byte(0), tree.Left.Right.Byte)
		assert.Equal(t, byte(97), tree.Left.Right.Left.Byte)
		assert.Equal(t, byte(0), tree.Left.Right.Right.Byte)
		assert.Equal(t, byte(33), tree.Left.Right.Right.Left.Byte)
		assert.Equal(t, byte(114), tree.Left.Right.Right.Right.Byte)

		assert.Equal(t, byte(0), tree.Right.Byte)
		assert.Equal(t, byte(101), tree.Right.Left.Byte)
		assert.Equal(t, byte(0), tree.Right.Right.Byte)
		assert.Equal(t, byte(111), tree.Right.Right.Left.Byte)
		assert.Equal(t, byte(112), tree.Right.Right.Right.Byte)
	})
}
