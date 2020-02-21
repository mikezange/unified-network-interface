package builder

import (
	"encoding/binary"
)

func (b *Builder) Float32LE(f float32) *Builder {
	return b.binaryWrite(f, binary.LittleEndian)
}

func (b *Builder) Float32BE(f float32) *Builder {
	return b.binaryWrite(f, binary.BigEndian)
}

func (b *Builder) Float64LE(f float64) *Builder {
	return b.binaryWrite(f, binary.LittleEndian)
}

func (b *Builder) Float64BE(f float64) *Builder {
	return b.binaryWrite(f, binary.BigEndian)
}