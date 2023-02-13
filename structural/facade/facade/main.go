package main

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{width: width, height: height,
		buffer: make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(b *Buffer) *Viewport {
	return &Viewport{buffer: b, offset: 0}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// a facade over buffers and viewports
type Console struct {
	Buffers   []*Buffer
	Viewports []*Viewport
	Offset    int
}

func NewConsole() *Console {
	b := NewBuffer(150, 200)
	v := NewViewport(b)
	return &Console{Buffers: []*Buffer{b}, Viewports: []*Viewport{v}}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.Viewports[0].GetCharacterAt(index)
}

func main() {

}
