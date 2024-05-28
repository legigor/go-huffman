package huffman

type freqTable map[byte]int

func (t freqTable) serialize() []byte {
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

func (t freqTable) deserialize(serialized []byte) {
	for i := 0; i < len(serialized); i += 5 {
		t[serialized[i]] = int(serialized[i+1])<<24 | int(serialized[i+2])<<16 | int(serialized[i+3])<<8 | int(serialized[i+4])
	}
}
