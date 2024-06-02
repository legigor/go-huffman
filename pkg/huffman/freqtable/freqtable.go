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
