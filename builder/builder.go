package builder

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Builder struct {
	reader io.Reader
	writer io.Writer
	buffer *bytes.Buffer
}

func New() *Builder {
	return &Builder{
		buffer: &bytes.Buffer{},
	}
}

func (b *Builder) WithReadWriter(rw io.ReadWriter) *Builder {
	b.reader = rw
	b.writer = rw
	return b
}

func (b *Builder) WithReader(r io.Reader) *Builder {
	b.reader = r
	return b
}

func (b *Builder) WithWriter(w io.Writer) *Builder {
	b.writer = w
	return b
}

func (b *Builder) ReadBuffer(buf []byte) (err error) {
	_, err = io.ReadFull(b.reader, buf)
	return
}

func (b *Builder) binaryWrite(data interface{}, order binary.ByteOrder) *Builder{
	_ = binary.Write(b.buffer, order, data)
	return b
}

func (b *Builder) binaryRead(out interface{}, order binary.ByteOrder) (interface{}, error) {
	err := binary.Read(b.reader, order, &out)
	return out, err
}

func (b *Builder) bufferWrite(data []byte) *Builder{
	b.buffer.Write(data)
	return b
}

func (b *Builder) bufferWriteString(s string) *Builder{
	b.buffer.WriteString(s)
	return b
}

func (b *Builder) bufferReadByte() (byte, error) {
	var buf [1]byte
	_, err := io.ReadFull(b.reader, buf[:])
	return buf[0], err
}

func (b *Builder) bufferReadBytes(size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := io.ReadFull(b.reader, buf)
	return buf, err
}

func (b *Builder) Make() []byte {
	return b.buffer.Bytes()
}