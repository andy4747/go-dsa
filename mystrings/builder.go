package mystrings

type Builder struct {
	buf []byte
}

func (b *Builder) WriteString(data string) (int, error) {
	b.buf = append(b.buf, data...)
	return len(data), nil
}

func (b *Builder) String() string {
	return string(b.buf)
}
