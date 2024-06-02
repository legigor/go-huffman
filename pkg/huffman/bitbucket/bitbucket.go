package bitbucket

type Bucket struct {
	bits     []byte
	shift    int
	position int
}

func NewBucket(bits []byte) *Bucket {
	return &Bucket{
		bits:     bits,
		shift:    0,
		position: 0,
	}
}

func (b *Bucket) Append(bits ...byte) {
	if len(b.bits) == 0 {
		b.bits = append(b.bits, 0)
	}

	for i := 0; i < len(bits); i++ {
		bit := bits[i]
		b.bits[b.position] = b.bits[b.position] | (bit << (7 - b.shift))
		b.shift++
		if b.shift == 8 {
			b.shift = 0
			b.position++
			b.bits = append(b.bits, 0)
		}
	}
}

func (b *Bucket) Bytes() []byte {
	return b.bits
}

type Iterator struct {
	bucket *Bucket

	position int
	shift    int
}

func (b *Bucket) Iterator() *Iterator {
	return &Iterator{
		bucket:   b,
		position: 0,
		shift:    0,
	}
}

func (i *Iterator) Read() byte {
	bit := i.bucket.bits[i.position] >> (7 - i.shift) & 1

	if i.shift == 7 {
		i.shift = 0
		i.position++
	} else {
		i.shift++
	}

	return bit
}

func (i *Iterator) HasData() bool {
	return i.position < len(i.bucket.bits)
}
