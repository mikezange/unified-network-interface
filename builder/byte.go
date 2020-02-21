package builder

import (
	"encoding/binary"
)

func (b *Builder) Byte(data byte) *Builder {
	return b.binaryWrite(data, binary.BigEndian)
}

func (b *Builder) Bytes(data []byte) *Builder {
	return b.bufferWrite(data)
}

func (b *Builder) ReadByte() (byte, error) {
	return b.bufferReadByte()
}

func (b *Builder) ReadBytes(len int) ([]byte, error) {
	return b.bufferReadBytes(len)
}