package builder

import (
	"encoding/binary"
)

func (b *Builder) IntLE(i int) *Builder {
	return b.binaryWrite(i, binary.LittleEndian)
}

func (b *Builder) IntBE(i int) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Int8(i int8) *Builder {
	return b.binaryWrite(byte(i), binary.BigEndian)
}

func (b *Builder) Int16LE(i int16) *Builder {
	return b.binaryWrite(i, binary.LittleEndian)
}

func (b *Builder) Int16BE(i int16) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Int32LE(i int32) *Builder {
	return b.binaryWrite(i, binary.LittleEndian)
}

func (b *Builder) Int32BE(i int32) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Int64LE(i int64) *Builder {
	return b.binaryWrite(i, binary.LittleEndian)
}

func (b *Builder) Int64BE(i int64) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Int(order binary.ByteOrder, i int64, size int) *Builder {
	data := make([]uint8, size)

	if order == binary.BigEndian {
		for x := 0; x < size; x++ {
			data[x] = byte(i >> byte((size-x-1)*8))
		}
	} else {
		for x := 0; x < size; x++ {
			data[x] = byte(i >> byte(x*8))
		}
	}

	return b.bufferWrite(data)
}