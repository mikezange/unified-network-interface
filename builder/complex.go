package builder

import (
	"encoding/binary"
)

func (b *Builder) Complex64LE(c complex64) *Builder {
	return b.binaryWrite(c, binary.LittleEndian)
}

func (b *Builder) Complex64BE(c complex64) *Builder {
	return b.binaryWrite(c, binary.BigEndian)
}

func (b *Builder) Complex128LE(c complex128) *Builder {
	return b.binaryWrite(c, binary.LittleEndian)
}

func (b *Builder) Complex128BE(c complex128) *Builder {
	return b.binaryWrite(c, binary.BigEndian)
}