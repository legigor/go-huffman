package freqtable

import "sort"

type Table map[byte]int

func Initialize(p []byte) Table {
	freq := Table{}
	for _, b := range p {
		freq[b]++
	}
	// Normalize
	var kv []keyValue
	for k, v := range freq {
		kv = append(kv, keyValue{Key: k, Value: v})
	}
	sort.Sort(byKeyValue(kv))
	for i := 0; i < len(kv); i++ {
		freq[kv[i].Key] = i + 1
	}
	return freq
}

type keyValue struct {
	Key   byte
	Value int
}

type byKeyValue []keyValue

func (a byKeyValue) Len() int      { return len(a) }
func (a byKeyValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byKeyValue) Less(i, j int) bool {
	if a[i].Value == a[j].Value {
		return a[i].Key < a[j].Key
	}

	return a[i].Value < a[j].Value
}

// Serialize TODO: Remove as we are serializing the tree
func (t Table) Serialize() []byte {
	buf := make([]byte, len(t)*5)
	pos := 0
	for k, v := range t {
		buf[pos] = k
		buf[pos+1] = byte(v >> 24)
		buf[pos+2] = byte(v >> 16)
		buf[pos+3] = byte(v >> 8)
		buf[pos+4] = byte(v)
		pos += 5
	}
	return buf
}

func (t Table) Deserialize(serialized []byte) {
	for i := 0; i < len(serialized); i += 5 {
		t[serialized[i]] = int(serialized[i+1])<<24 | int(serialized[i+2])<<16 | int(serialized[i+3])<<8 | int(serialized[i+4])
	}
}
