package freqtable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_serialize_deserialize(t *testing.T) {
	freq := Table{'a': 1, 'b': 2, 'c': 3}

	serialized := freq.Serialize()

	deserialized := Table{}
	deserialized.Deserialize(serialized)

	assert.Equal(t, freq, deserialized)
}

func Test_normalization(t *testing.T) {

	t.Run("different frequencies", func(t *testing.T) {
		freq1 := Initialize([]byte("abbccc"))
		assert.Equal(t, 1, freq1['a'])
		assert.Equal(t, 2, freq1['b'])
		assert.Equal(t, 3, freq1['c'])
	})

	t.Run("same frequencies", func(t *testing.T) {
		freq1 := Initialize([]byte("bb_cc_aa_e_d_fff_gggg"))
		testData := map[string]int{
			"a": 3,
			"b": 4,
			"c": 5,
			"d": 1,
			"e": 2,
			"f": 6,
			"g": 7,
			"_": 8,
		}
		for k, v := range testData {
			t.Run(fmt.Sprintf("'%s'=%d", k, v), func(t *testing.T) {
				assert.Equal(t, v, freq1[k[0]])
			})
		}
	})
}
