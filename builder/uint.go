package builder

import (
	"encoding/binary"
)

func (b *Builder) UintLE(i uint) *Builder {
	return b.binaryWrite(i, binary.LittleEndian)
}

func (b *Builder) UintBE(i uint) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Uint8(i uint8) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Uint16LE(i uint16) *Builder {
	return b.binaryWrite(i, binary.LittleEndian)
}

func (b *Builder) Uint16BE(i uint16) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Uint32LE(i uint32) *Builder {
	return b.binaryWrite(i, binary.LittleEndian)
}

func (b *Builder) Uint32BE(i uint32) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Uint64LE(i uint64) *Builder {
	return b.binaryWrite(i, binary.LittleEndian)
}

func (b *Builder) Uint64BE(i uint64) *Builder {
	return b.binaryWrite(i, binary.BigEndian)
}

func (b *Builder) Uint(order binary.ByteOrder, i uint64, size int) *Builder {
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

func (b *Builder) ReadUint8() (byte, error) {
	return b.bufferReadByte()
}

func (b *Builder) ReadUintBE() (i uint, err error) {
	val, err := b.binaryRead(i, binary.BigEndian)
	return val.(uint), err
}

func (b *Builder) ReadUintLE() (i uint, err error) {
	val, err := b.binaryRead(i, binary.LittleEndian)
	i = val.(uint)
	return
}