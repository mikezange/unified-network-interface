package builder

func (b *Builder) String(s string) *Builder {
	return b.bufferWriteString(s)
}

func (b *Builder) ReadString(len int) (string, error) {
	buf, err := b.bufferReadBytes(len)
	return string(buf), err
}